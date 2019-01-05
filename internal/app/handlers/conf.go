package handlers

import (
	"errors"

	"github.com/ltheinrich/gorum/pkg/config"
)

// Conf handler
func Conf(request map[string]interface{}, username string, auth bool) interface{} {
	// check if confkeys are provided
	confkeys := GetStringArray(request, "confkeys")
	if confkeys == nil {
		// return not provided
		return errors.New("400")
	}

	// map to write
	confvalues := map[string]interface{}{}

	// loop through confkeys
	for _, confkey := range confkeys {
		// set confvalue in map
		confvalues[confkey] = config.Get("public", confkey)
	}

	// return map
	return confvalues
}
