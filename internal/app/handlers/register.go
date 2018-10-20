package handlers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lheinrichde/gorum/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// Register handler
func Register(request map[string]interface{}, username string) interface{} {
	var err error

	// check if username and password are provided
	username, mail, password := GetString(request, "username"), GetString(request, "mail"), GetString(request, "password")
	if username == "" || mail == "" || password == "" {
		// return not provided
		return errors.New("400")
	}

	// query db
	var id int
	err = db.DB.QueryRow("SELECT id FROM users WHERE username = $1 OR mail = $2;", username, mail).Scan(&id)
	if err == sql.ErrNoRows {
		// not exists
		var passwordHash []byte
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+1)
		if err != nil {
			// return error
			return err
		}

		// insert into database
		_, err = db.DB.Exec("INSERT INTO users (username, passwordhash, mail, registered) VALUES ($1, $2, $3, $4);", username, string(passwordHash), mail, time.Now().Format("2006-01-02T15:04:05"))
		if err != nil {
			// return error
			return err
		}

		// registered
		return map[string]interface{}{"done": true}
	} else if err != nil {
		// return error
		return err
	}

	// username exists
	return map[string]interface{}{"done": false}
}
