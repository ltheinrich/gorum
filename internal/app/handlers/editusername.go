package handlers

import (
	"database/sql"
	"errors"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// EditUsername handler
func EditUsername(data HandlerData) interface{} {
	var err error

	// authenticate
	if !data.Authenticated {
		// not authenticated
		return errors.New("403")
	}

	// check if new username and password is provided
	newUsername := data.Request.GetString("newUsername")
	if newUsername == "" || len(newUsername) > 32 {
		// both not provided
		return errors.New("400")
	}

	// check if new username already exists
	var trash int
	err = db.DB.QueryRow("SELECT id from users WHERE username = $1;", newUsername).Scan(&trash)
	if err != sql.ErrNoRows {
		// username already exists
		return errors.New("usernameExists")
	} else if err != nil && err != sql.ErrNoRows {
		// return error
		return err
	}

	// update username
	_, err = db.DB.Exec("UPDATE users SET username = $1 WHERE username = $2;", newUsername, data.Username)
	if err != nil {
		// return error
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
