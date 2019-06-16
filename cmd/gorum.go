package cmd

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dchest/captcha"

	"github.com/ltheinrich/gorum/internal/app/handlers"
	"github.com/ltheinrich/gorum/internal/pkg/assets"
	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

var (
	// Version string generated using Makefile
	Version string

	// BuildTime string generated using Makefile
	BuildTime string
)

// Init startup
func Init() (err error) {
	// load config
	if err = loadConfig(); err != nil {
		return err
	}

	// load language
	if err = loadLanguage(); err != nil {
		return err
	}

	// connect to postgresql database
	if err = connectDB(); err != nil {
		return err
	}

	// run setup query
	if err = setupDB(); err != nil {
		return err
	}

	// register handlers
	handle()

	// https listen
	return listen()
}

// register handlers
func handle() {
	// web files (Angular)
	http.HandleFunc("/", handlers.Web)

	// data files
	http.HandleFunc("/data/", handlers.Data)

	// custom handlers
	http.HandleFunc("/uploadavatar", handlers.UploadAvatar)
	http.Handle("/captcha/", captcha.Server(240, 80))

	// register all handlers in map
	for url, handler := range handlers.Handlers {
		RegisterHandler(url, handler)
	}
}

// RegisterHandler add handler
func RegisterHandler(url string, handler func(data handlers.HandlerData) interface{}) {
	http.HandleFunc("/api/"+url, handlers.GenerateHandler(handler))
}

// load config template and overwrite with custom
func loadConfig() (err error) {
	// load template config
	templateConfig := assets.MustAsset("config.tpl.json")
	if err = config.ProcessConfig(templateConfig); err != nil {
		return err
	}

	// load custom config
	return config.LoadConfig("config.json")
}

// load language file
func loadLanguage() (err error) {
	// load language file and set
	languageFile := assets.MustAsset("language.json")

	// unmarshal to map and check for error
	var languages map[string]map[string]string
	err = json.Unmarshal(languageFile, &languages)
	if err != nil {
		return err
	}

	// loop through language maps
	for language := range languages {
		// check if language not German
		if language != "de" {
			// loop through language keys
			for key, value := range languages["de"] {
				// check if language variable exists
				if _, exists := languages[language][key]; !exists {
					// complete with soft fallback language (German)
					languages[language][key] = value
				}
			}
		}

		// marshal map and check for error
		var languageData []byte
		languageData, err = json.Marshal(languages[language])
		if err != nil {
			return err
		}

		// set language
		handlers.Languages[language] = languageData
	}

	return nil
}

// connect to PostgreSQL database
func connectDB() (err error) {
	// define login variables
	host := config.Get("postgresql", "host")
	port := config.Get("postgresql", "port")
	ssl := config.Get("postgresql", "ssl")
	database := config.Get("postgresql", "database")
	username := config.Get("postgresql", "username")
	password := config.Get("postgresql", "password")

	// connect and return error
	return db.Connect(host, port, ssl, database, username, password)
}

// run setup query
func setupDB() (err error) {
	// get file
	query := assets.MustAsset("setup.sql")

	// return error
	_, err = db.DB.Exec(string(query))
	return err
}

// listen to address
func listen() (err error) {
	// define http(s) variable
	address := config.Get("https", "address")
	certificate := config.Get("https", "certificate")
	key := config.Get("https", "key")

	// address as usable url
	url := address
	if strings.HasPrefix(address, ":") {
		url = "localhost" + address
	}

	// check if certicate and key file provided
	if certificate == "" || key == "" {
		// http server
		log.Printf("Webserver listening at http://%v/\n", url)
		server := &http.Server{Addr: address}
		return server.ListenAndServe()
	}

	// enable TLS 1.3
	if config.GetBool("https", "tls13") {
		os.Setenv("GODEBUG", "tls13=1")
		log.Println("Explicitly enabled TLS 1.3 for the https web server")
	}

	// https/tls server
	log.Printf("Webserver listening at https://%v/\n", url)
	server := &http.Server{Addr: address,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
			CurvePreferences: []tls.CurveID{
				tls.X25519,
				tls.CurveP256,
				tls.CurveP384,
				tls.CurveP521,
			},
		}}
	return server.ListenAndServeTLS(certificate, key)
}
