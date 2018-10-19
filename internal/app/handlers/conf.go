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

	// check if confkey is provided
	confkey := GetString(request, "confkey")
	if confkey == "" {
		Code(w, "400")
		return
	}

	// query db
	var confvalue string
	err = db.DB.QueryRow("SELECT confvalue FROM config WHERE confkey = $1;", confkey).Scan(&confvalue)
	if err != nil {
		// error
		Error(w, err)
		return
	}

	// write
	Write(w, map[string]interface{}{confkey: confvalue})
}
