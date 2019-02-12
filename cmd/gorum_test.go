package cmd

import (
	"errors"
	"testing"
	"time"

	"github.com/ltheinrich/gorum/internal/pkg/db"

	"github.com/ltheinrich/gorum/internal/app/handlers"

	"github.com/ltheinrich/gorum/internal/pkg/config"
)

func TestLoadConfig(t *testing.T) {
	// call loadConfig and get error
	err := loadConfig()

	// check if error occurred
	if err != nil {
		// failed
		t.Errorf("Could not load config, %v\n", err)
	} else if config.Get("postgresql", "host") == "" {
		// check if PostgreSQL host is not set, failed
		t.Error("Could not load config, PostgreSQL host is not set")
	}
}

func TestLoadLanguage(t *testing.T) {
	// call loadLanguage
	err := loadLanguage()
	if err != nil {
		t.Errorf("Could not load language, %v\n", err)
		return
	}

	// check if German language (soft fallback) is set
	if handlers.Languages["de"] == nil {
		// failed
		t.Error("Could not load language, German language (soft fallback) is not set")
	}
}

func TestConnectDB(t *testing.T) {
	// load config and get error
	err := loadConfig()
	if err != nil {
		// failed
		t.Errorf("Could not connect to database, config could not be loaded: %v\n", err)
		return
	}

	// call connectDB and get error
	err = connectDB()

	// check if error occurred
	if err != nil {
		// failed
		t.Errorf("Could not connect to database, %v\n", err)
	} else if db.DB == nil {
		// check if database variable is nil, failed
		t.Error("Could not connect to database, database variable is nil")
	}
}

func TestSetupDB(t *testing.T) {
	// load config and get error
	err := loadConfig()
	if err != nil {
		// failed
		t.Errorf("Could not setup database, config could not be loaded: %v\n", err)
		return
	}

	// connect to database and get error
	err = connectDB()
	if err != nil {
		// failed
		t.Errorf("Could not setup database, database connection could not be opened: %v\n", err)
		return
	}

	// call setupDB
	err = setupDB()
	if err != nil {
		// failed
		t.Errorf("Could not setup database, %v\n", err)
	}
}

func TestListen(t *testing.T) {
	// load config and get error
	err := loadConfig()
	if err != nil {
		// failed
		t.Errorf("Could not listen, config could not be loaded: %v\n", err)
		return
	}

	// make channel for error
	fail := make(chan error)

	// gofunction to call listen
	go func() {
		// call listen and get error
		err := listen()
		if err == nil {
			// no error, but malfunction
			err = errors.New("listen function ran through")
		}

		// listen ran through, send error
		fail <- err
	}()

	// gofunction for timeout
	go func() {
		// let listen function be running for 3 seconds
		time.Sleep(3 * time.Second)

		// listen function did not finish
		fail <- nil
	}()

	// check whether failed
	if err := <-fail; err != nil {
		// print error message
		t.Errorf("Could not listen, %v\n", err)
	}
}
