package handlers

import (
	"errors"

	"github.com/lheinrichde/gorum/pkg/db"
)

// EditUsername handler
func EditUsername(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// authenticate
	if !auth {
		// not authenticated
		return errors.New("403")
	}

	// check if new username and password is provided
	newUsername := GetString(request, "newUsername")
	if newUsername == "" {
		// both not provided
		return errors.New("400")
	}

	// update username
	_, err = db.DB.Exec("UPDATE users SET username = $1 WHERE username = $2;", newUsername, username)

	// check error
	if err != nil {
		// return error
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
