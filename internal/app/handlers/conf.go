package handlers

import (
	"encoding/json"
	"log"

	"github.com/ltheinrich/gorum/internal/pkg/config"
)

var (
	configMap []byte
)

// Conf handler
func Conf(data HandlerData) interface{} {
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
			// print and return error
			log.Println(err)
			return err
		}
	}

	// return public config map bytes
	return configMap
}
