package handlers

import (
	"net/http"
)

var (
	// Language file
	Language []byte
)

// Lang handler
func Lang(w http.ResponseWriter, r *http.Request) {
	Header(w)

	// write
	w.Write(Language)
}
