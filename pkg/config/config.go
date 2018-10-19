package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	// configuration map
	config map[string]map[string]string

	// configuration file name
	file string
)

// Get configuration value
func Get(parent, child string) string {
	return config[parent][child]
}

// Set configuration value and save
func Set(parent, child, value string) error {
	var err error

	// set value
	config[parent][child] = value

	// marshal
	var data []byte
	data, err = json.Marshal(config)
	if err != nil {
		return err
	}

	// write to file
	err = ioutil.WriteFile(file, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// LoadConfig load json config
func LoadConfig(fileName string) error {
	var err error

	// set file name
	file = fileName

	// read file
	var data []byte
	data, err = ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		data = []byte("{}")
	} else if err != nil {
		return err
	}

	// unmarshal data to temporary config
	var tempConfig map[string]map[string]string
	err = json.Unmarshal(data, &tempConfig)
	if err != nil {
		return err
	}

	// update sub configs
	for name, subConfig := range tempConfig {
		// check if config exists
		if config == nil {
			// create config
			config = map[string]map[string]string{}
		}

		// check if sub config exists
		if config[name] == nil {
			// create sub config
			config[name] = subConfig
		} else {
			// update values
			for key, value := range subConfig {
				config[name][key] = value
			}
		}
	}

	return nil
}
