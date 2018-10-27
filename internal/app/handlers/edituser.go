package handlers

import (
	"errors"

	"github.com/lheinrichde/gorum/pkg/db"
)

// EditUser handler
func EditUser(request map[string]interface{}, username string) interface{} {
	var err error

	// check if new user data provided
	newUsername, newAvatar := GetString(request, "newUsername"), GetString(request, "newAvatar")
	if newUsername == "" && newAvatar == "" {
		// both not provided
		return errors.New("400")
	}

	// check what provided
	if newUsername != "" {
		// update username
		_, err = db.DB.Exec("UPDATE users SET username = $1 WHERE username = $2;", newUsername, username)
	} else if newAvatar != "" {
		// update avatar
		_, err = db.DB.Exec("UPDATE users SET avatar = $1 WHERE username = $2;", newAvatar, username)
	}

	// check error
	if err != nil {
		// return error
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
