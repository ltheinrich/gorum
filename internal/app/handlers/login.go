package handlers

import (
	"database/sql"
	"net/http"

	"github.com/lheinrichde/golib/pkg/handler"

	"golang.org/x/crypto/bcrypt"

	"github.com/lheinrichde/golib/pkg/db"
)

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	Header(w)
	request, err := Read(r.Body, r.ContentLength)
	if err != nil {
		Error(w, err)
		return
	}

	// check if username and password are provided
	username, password := handler.GetString(request, "username"), handler.GetString(request, "password")
	if username == "" || password == "" {
		ErrorWrite(w, "400")
		return
	}

	// query db
	var passwordHash string
	err = db.DB.QueryRow("SELECT passwordhash FROM users WHERE username = $1 OR mail = $1;", username).Scan(&passwordHash)
	if err == sql.ErrNoRows {
		// not exists
		ErrorWrite(w, "403")
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
