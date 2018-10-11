package gorum

import (
	"github.com/lheinrichde/golib/pkg/config"
)

// Init startup
func Init() error {
	if err := loadConfig(); err != nil {
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
