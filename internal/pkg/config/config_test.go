package config

import (
	"testing"
)

func TestProcessConfig(t *testing.T) {
	// process example config
	err := ProcessConfig([]byte(`{"Hallo": {"Welt": "Test"}}`))
	if err != nil {
		// failed
		t.Errorf("Could not process config, %v\n", err)
		return
	}

	// check if config value is correct
	if Get("Hallo", "Welt") != "Test" {
		// failed
		t.Error("Could not process config, example config value is wrong")
	}
}

func TestLoadConfig(t *testing.T) {
	// load example config
	err := LoadConfig("../../../assets/config.tpl.json")
	if err != nil {
		// failed
		t.Errorf("Could not load config, %v\n", err)
		return
	}

	// check if PostgreSQL host value is localhost
	if Get("postgresql", "host") != "localhost" {
		// failed
		t.Error("Could not process config, example config value is wrong")
	}
}

func TestGet(t *testing.T) {
	// process example config
	err := ProcessConfig([]byte(`{"Hallo": {"Welt": "Test"}}`))
	if err != nil {
		// failed
		t.Errorf("Could not get config value, could not process example config: %v\n", err)
		return
	}

	// check if config value is correct
	if Get("Hallo", "Welt") != "Test" {
		// failed
		t.Error("Could get config value, example config value is wrong")
	}
}

func TestGetInt(t *testing.T) {
	// process example config
	err := ProcessConfig([]byte(`{"Hallo": {"Welt": "1"}}`))
	if err != nil {
		// failed
		t.Errorf("Could not get integer config value, could not process example config: %v\n", err)
		return
	}

	// check if config value is correct
	if GetInt("Hallo", "Welt") != 1 {
		// failed
		t.Error("Could get integer config value, example config value is wrong")
	}
}

func TestSub(t *testing.T) {
	// process example config
	err := ProcessConfig([]byte(`{"Hallo": {"Welt": "Test"}}`))
	if err != nil {
		// failed
		t.Errorf("Could not get sub config, could not process example config: %v\n", err)
		return
	}

	// check if config value is correct
	if value, _ := Sub("Hallo")["Welt"]; value != "Test" {
		// failed
		t.Error("Could get sub config, example config value is wrong")
	}
}
