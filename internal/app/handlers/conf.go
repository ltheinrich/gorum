package handlers

import (
	"encoding/json"

	"github.com/ltheinrich/gorum/pkg/config"
)

var (
	configMap []byte
)

// Conf handler
func Conf(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check if config map is loaded
	if configMap == nil {
		// cast map
		confMap := map[string]interface{}{}
		for key, value := range config.Sub("public") {
			confMap[key] = value
		}

		// marshal map
		configMap, err = json.Marshal(confMap)
		if err != nil {
			return err
		}
	}

	// return public config map bytes
	return configMap
}
