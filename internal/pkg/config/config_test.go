package config

import "testing"

func TestProcessConfig(t *testing.T) {
	// ../../../assets/config.tpl.json
	err := ProcessConfig(`{"Hallo": "Welt"}`)
}
