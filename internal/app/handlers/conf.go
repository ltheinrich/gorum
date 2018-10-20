package handlers

import (
	"net/http"

	"github.com/lheinrichde/gorum/pkg/db"
)

// Conf handler
func Conf(w http.ResponseWriter, r *http.Request) {
	var err error
	Header(w)
	request := Read(r.Body, r.ContentLength)

	// check if confkeys are provided
	confkeys := GetStringArray(request, "confkeys")
	if confkeys == nil {
		Code(w, "400")
		return
	}

	// map to write
	confvalues := map[string]interface{}{}

	// loop through confkeys
	for _, confkey := range confkeys {
		// query db
		var confvalue string
		err = db.DB.QueryRow("SELECT confvalue FROM config WHERE confkey = $1;", confkey).Scan(&confvalue)
		if err != nil {
			// error
			Error(w, err)
			return
		}

		// set confvalue in map
		confvalues[confkey] = confvalue
	}

	// write map
	Write(w, confvalues)
}
