package gorum

import (
	"fmt"
	"net/http"

	"github.com/ltheinrich/captcha"

	"github.com/ltheinrich/gorum/internal/app/handlers"
	"github.com/ltheinrich/gorum/internal/pkg/assets"
	"github.com/ltheinrich/gorum/pkg/config"
	"github.com/ltheinrich/gorum/pkg/db"
)

// Init startup
func Init() error {
	// load config
	if err := loadConfig(); err != nil {
		return err
	}

	// load language
	loadLanguage()

	// connect to postgresql database
	if err := connectDB(); err != nil {
		return err
	}

	// run setup query
	if err := setupDB(); err != nil {
		return err
	}

	// register handlers
	handle()
	fmt.Println("Gorum (c) 2018 Lennart Heinrich")

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
func RegisterHandler(url string, handler func(request map[string]interface{}, username string, auth bool) interface{}) {
	http.HandleFunc("/api/"+url, handlers.GenerateHandler(handler))
}

// load config template and overwrite with custom
func loadConfig() error {
	// load template config
	templateConfig := assets.MustAsset("config.tpl.json")
	if err := config.ProcessConfig(templateConfig); err != nil {
		return err
	}

	// load custom config
	return config.LoadConfig("config.json")
}

// load language file
func loadLanguage() {
	// load language file and set
	language := assets.MustAsset("language.json")
	handlers.Language = language
}

// connect to postgresql database
func connectDB() error {
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
func setupDB() error {
	var err error

	// get file
	query := assets.MustAsset("setup.sql")

	// return error
	_, err = db.DB.Exec(string(query))
	return err
}

// listen to address
func listen() error {
	// define http(s) variable
	address := config.Get("https", "address")
	certificate := config.Get("https", "certificate")
	key := config.Get("https", "key")

	// check if certicate and key file provided
	if certificate == "" || key == "" {
		// http server
		return http.ListenAndServe(address, nil)
	}

	// https/tls server
	return http.ListenAndServeTLS(address, certificate, key, nil)
}
