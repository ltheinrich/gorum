package handlers

import (
	"database/sql"
	"errors"

	"github.com/lheinrichde/gorum/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// Login handler
func Login(request map[string]interface{}, username string) interface{} {
	var err error

	// check if username and password are provided
	username, password := GetString(request, "username"), GetString(request, "password")
	if username == "" || password == "" {
		// return not provided
		return errors.New("400")
	}

	// query db
	var passwordHash string
	err = db.DB.QueryRow("SELECT passwordhash FROM users WHERE username = $1 OR mail = $1;", username).Scan(&passwordHash)
	if err == sql.ErrNoRows {
		// not exists, but due to security return invalid
		return errors.New("403")
	} else if err != nil {
		// error
		return err
	}

	// compare passwords and write
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return map[string]interface{}{"valid": err == nil}
}
