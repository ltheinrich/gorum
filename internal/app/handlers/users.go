package handlers

import (
	"database/sql"
	"strconv"

	"github.com/lheinrichde/gorum/pkg/db"
)

// Users handler
func Users(request map[string]interface{}, username string) interface{} {
	var err error

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT id, username, registered, avatar FROM users;")
	if err != nil {
		// return error
		return err
	}

	// users map to write
	users := map[string]interface{}{}

	// loop through users
	for rows.Next() {
		// scan
		var id int
		var username, registered, avatar string
		err = rows.Scan(&id, &username, &registered, &avatar)
		if err != nil {
			// return error
			return err
		}

		// user map to append
		user := map[string]interface{}{}
		user["username"] = username
		user["registered"] = registered
		user["avatar"] = avatar

		// append user to users map
		users[strconv.Itoa(id)] = user
	}

	// write map
	return users
}
