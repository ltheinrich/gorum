package handlers

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/lheinrichde/golib/pkg/handler"

	"github.com/lheinrichde/golib/pkg/db"
)

// Register handler
func Register(w http.ResponseWriter, r *http.Request) {
	Header(w)
	request, err := Read(r.Body, r.ContentLength)
	if err != nil {
		Error(w, err)
		return
	}

	// check if username and password are provided
	username, mail, password := handler.GetString(request, "username"), handler.GetString(request, "mail"), handler.GetString(request, "password")
	if username == "" || mail == "" || password == "" {
		ErrorWrite(w, "400")
	}

	// query db
	var id int
	err = db.DB.QueryRow("SELECT id FROM users WHERE username = $1 OR mail = $2;", username, mail).Scan(&id)
	if err == sql.ErrNoRows {
		// not exists
		var passwordHash []byte
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+1)
		if err != nil {
			Error(w, err)
			return
		}

		// insert into database
		_, err = db.DB.Exec("INSERT INTO users (username, passwordhash, mail) VALUES ($1, $2, $3);", username, string(passwordHash), mail)
		if err != nil {
			Error(w, err)
			return
		}

		// registered
		Write(w, map[string]interface{}{"done": true})
		return
	} else if err != nil {
		// error
		Error(w, err)
		return
	}

	// username exists
	Write(w, map[string]interface{}{"done": false})
}
