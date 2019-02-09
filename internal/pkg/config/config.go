package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	// configuration map
	config map[string]map[string]string

	// configuration file name
	file string
)

// Sub configuration map
func Sub(parent string) map[string]string {
	// get sub configuration map
	sub := config[parent]

	// check if nil
	if sub == nil {
		// set to empty map
		sub = map[string]string{}
	}

	// return map
	return sub
}

// Get configuration value
func Get(parent, child string) string {
	return Sub(parent)[child]
}

// GetInt configuration value as integer
func GetInt(parent, child string) int {
	// get and parse
	num, err := strconv.Atoi(Get(parent, child))
	if err != nil {
		// not an integer, return zero
		return 0
	}

	// return integer value
	return num
}

// Set configuration value and save
func Set(parent, child, value string) error {
	var err error

	// check if map is nil
	if config[parent] == nil {
		config[parent] = map[string]string{}
	}

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

	// empty json file if not exists
	if os.IsNotExist(err) {
		data = []byte("{}")
	} else if err != nil {
		return err
	}

	// process read file
	return ProcessConfig(data)
}

// ProcessConfig load configuration from []byte
func ProcessConfig(data []byte) error {
	var err error

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

	// no error
	return nil
}
