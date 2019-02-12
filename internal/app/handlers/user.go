package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ltheinrich/gorum/internal/pkg/config"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// User handler
func User(data HandlerData) interface{} {
	var err error

	// check if user data provided
	userID := data.Request.GetInt("userID")
	if userID == 0 && data.Username == "" {
		// both not provided
		return errors.New("400")
	}

	// define variables
	var queryID int
	var queryUsername, registered string

	// check what provided
	if userID == 0 {
		// query current user
		err = db.DB.QueryRow("SELECT id, username, registered FROM users WHERE username = $1;", data.Username).Scan(&queryID, &queryUsername, &registered)
	} else {
		// query user by id
		err = db.DB.QueryRow("SELECT id, username, registered FROM users WHERE id = $1;", userID).Scan(&queryID, &queryUsername, &registered)
	}

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// print and return error
		log.Println(err)
		return err
	}

	// user map to write
	user := map[string]interface{}{}
	user["id"] = queryID
	user["username"] = queryUsername
	user["registered"] = registered

	// add avatar
	avatarPath := fmt.Sprintf("%s/%v.png", config.Get("data", "avatar"), queryID)
	_, err = os.Open(avatarPath)
	if os.IsNotExist(err) {
		user["avatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
	} else {
		user["avatar"] = avatarPath
	}

	// write map
	return user
}
