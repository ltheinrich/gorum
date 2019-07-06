package handlers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/NathanNr/gorum/internal/pkg/db"
)

// EditPassword handler
func EditPassword(data HandlerData) interface{} {
	var err error

	// authenticate
	password := data.Request.GetString("password")
	if password == "" || !data.Authenticated || !login(data.Username, password) {
		// not authenticated
		return errors.New("403")
	}

	// check if new password provided
	newPassword := data.Request.GetString("newPassword")
	if newPassword == "" {
		// not provided
		return errors.New("400")
	}

	// generate new password hash
	var passwordHash []byte
	passwordHash, err = bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost+1)
	if err != nil {
		// return error
		return err
	}

	// update password
	_, err = db.DB.Exec("UPDATE users SET passwordhash = $1 WHERE username = $2;", passwordHash, data.Username)
	if err != nil {
		// return error
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
