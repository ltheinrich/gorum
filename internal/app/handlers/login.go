package handlers

import (
	"database/sql"
	"net/http"

	"github.com/lheinrichde/gorum/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	Header(w)
	request := Read(r.Body, r.ContentLength)

	// check if username and password are provided
	username, password := GetString(request, "username"), GetString(request, "password")
	if username == "" || password == "" {
		Code(w, "400")
		return
	}

	// query db
	var passwordHash string
	err = db.DB.QueryRow("SELECT passwordhash FROM users WHERE username = $1 OR mail = $1;", username).Scan(&passwordHash)
	if err == sql.ErrNoRows {
		// not exists
		Code(w, "403")
		return
	} else if err != nil {
		// error
		Error(w, err)
		return
	}

	// compare passwords and write
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	Write(w, map[string]interface{}{"valid": err == nil})
}
