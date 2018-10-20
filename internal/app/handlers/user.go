package handlers

import (
	"errors"

	"github.com/lheinrichde/gorum/pkg/db"
)

// User handler
func User(request map[string]interface{}, username string) interface{} {
	var err error

	// check if user id provided
	userID := GetInt(request, "userID")
	if userID == 0 {
		// return not provided
		return errors.New("400")
	}

	// query db
	var queryUsername, registered, avatar string
	err = db.DB.QueryRow("SELECT username, registered, avatar FROM users WHERE id = $1;", userID).Scan(&queryUsername, &registered, &avatar)
	if err != nil {
		// return error
		return err
	}

	// user map to write
	user := map[string]interface{}{}
	user["username"] = queryUsername
	user["registered"] = registered
	user["avatar"] = avatar

	// write map
	return user
}
