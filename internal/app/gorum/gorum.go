package gorum

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lheinrichde/golib/pkg/config"
	"github.com/lheinrichde/golib/pkg/db"
	"github.com/lheinrichde/gorum/internal/app/handlers"
)

// Init startup
func Init() error {
	// load config
	if err := loadConfig(); err != nil {
		return err
	}

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
	if err := listen(); err != nil {
		return err
	}

	return nil
}

// register handlers
func handle() {
	// web/dist files (Angular)
	http.Handle("/", http.FileServer(http.Dir("web/dist/gorum")))

	// register all handlers in map
	for url, h := range handlers.Handlers {
		http.HandleFunc("/api/"+url, h)
	}
}

// load config template and overwrite with custom
func loadConfig() error {
	// load template config
	if err := config.LoadConfig("assets/config.tpl.json"); err != nil {
		return err
	}

	// load custom config
	if err := config.LoadConfig("config.json"); err != nil {
		return err
	}

	return nil
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
	var query []byte
	query, err = ioutil.ReadFile("assets/setup.sql")
	if err != nil {
		return err
	}

	// return error
	_, err = db.DB.Exec(string(query))
	return err
}

// listen to address (https)
func listen() error {
	// define https variable
	address := config.Get("https", "address")
	certificate := config.Get("https", "certificate")
	key := config.Get("https", "key")

	// return error
	return http.ListenAndServeTLS(address, certificate, key, nil)
}
