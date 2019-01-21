package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// Users handler
func Users(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT id, username, registered FROM users;")
	if err != nil {
		// print and return error
		log.Println(err)
		return err
	}
	defer rows.Close()

	// users map to write
	users := map[string]interface{}{}

	// loop through users
	for rows.Next() {
		// scan
		var id int
		var queryUsername, registered string
		err = rows.Scan(&id, &queryUsername, &registered)
		if err != nil {
			// print and return error
			log.Println(err)
			return err
		}

		// user map to append
		user := map[string]interface{}{}
		user["username"] = queryUsername
		user["registered"] = registered

		// add avatar
		avatarPath := fmt.Sprintf("%s/%v.png", config.Get("data", "avatar"), id)
		_, err = os.Open(avatarPath)
		if os.IsNotExist(err) {
			user["avatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
		} else {
			user["avatar"] = avatarPath
		}

		// append user to users map
		users[strconv.Itoa(id)] = user
	}

	// write map
	return users
}
