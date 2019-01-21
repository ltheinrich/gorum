package handlers

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// EditPassword handler
func EditPassword(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// authenticate
	if !auth {
		// not authenticated
		return errors.New("403")
	}

	// check if new password provided
	newPassword := GetString(request, "newPassword")
	if newPassword == "" {
		// not provided
		return errors.New("400")
	}

	// generate new password hash
	var passwordHash []byte
	passwordHash, err = bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost+1)
	if err != nil {
		// print and return error
		log.Println(err)
		return err
	}

	// update password
	_, err = db.DB.Exec("UPDATE users SET passwordhash = $1 WHERE username = $2;", passwordHash, username)
	if err != nil {
		// print and return error
		log.Println(err)
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
