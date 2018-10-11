package gorum

import (
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

	return nil
}

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

func connectDB() error {
	host := config.Get("postgresql", "host")
	port := config.Get("postgresql", "port")
	ssl := config.Get("postgresql", "ssl")
	database := config.Get("postgresql", "database")
	username := config.Get("postgresql", "username")
	password := config.Get("postgresql", "password")

	return db.Connect(host, port, ssl, database, username, password)
}
