package handlers

import (
	"errors"

	"github.com/lheinrichde/gorum/pkg/db"
)

// Conf handler
func Conf(request map[string]interface{}, username string) interface{} {
	var err error

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
		// query db
		var confvalue string
		err = db.DB.QueryRow("SELECT confvalue FROM config WHERE confkey = $1;", confkey).Scan(&confvalue)
		if err != nil {
			// return error
			return err
		}

		// set confvalue in map
		confvalues[confkey] = confvalue
	}

	// return map
	return confvalues
}
