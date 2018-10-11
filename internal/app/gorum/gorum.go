package gorum

import (
	"fmt"
	"net/http"

	"github.com/lheinrichde/golib/pkg/config"
	"github.com/lheinrichde/golib/pkg/db"
)

// Init startup
func Init() error {
	if err := loadConfig(); err != nil {
		return err
	}

	if err := connectDB(); err != nil {
		return err
	}

	handle()
	fmt.Println("Gorum (c) 2018 Lennart Heinrich")

	if err := listen(); err != nil {
		return err
	}

	return nil
}

// register handlers
func handle() {
	// web/dist files (Angular)
	http.Handle("/", http.FileServer(http.Dir("web/dist/gorum")))
}

// load config template and overwrite with custom
func loadConfig() error {
	// load template config
	if err := config.LoadConfig("assets/config.tpl.json"); err != nil {
		return err
	}

	// load custom config
	if err := config.LoadConfig("assets/config.json"); err != nil {
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

// listen to address (https)
func listen() error {
	// define https variable
	address := config.Get("https", "address")
	certificate := config.Get("https", "certificate")
	key := config.Get("https", "key")

	return http.ListenAndServeTLS(address, certificate, key, nil)
}
