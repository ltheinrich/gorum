package handlers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/lheinrichde/gorum/pkg/db"
)

// ChangePassword handler
func ChangePassword(request map[string]interface{}, username string) interface{} {
	var err error

	// check if new- and password provided
	newPassword, password := GetString(request, "newPassword"), GetString(request, "password")
	if newPassword == "" || password == "" {
		// both not provided
		return errors.New("400")
	}

	// check login
	if !login(username, password) {
		// invalid login
		return errors.New("403")
	}

	// generate new password hash
	var passwordHash []byte
	passwordHash, err = bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost+1)
	if err != nil {
		return err
	}

	// update password
	_, err = db.DB.Exec("UPDATE users SET passwordhash = $1 WHERE username = $2;", passwordHash, username)
	if err != nil {
		// return error
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
