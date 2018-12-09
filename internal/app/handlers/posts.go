package handlers

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/lheinrichde/gorum/pkg/db"
)

// Posts handler
func Posts(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// get thread id and check if provided
	threadID := GetInt(request, "boardID")
	if threadID == 0 {
		// no thread id provided
		return errors.New("400")
	}

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT id, author, created, content WHERE thread = $1;`, threadID)
	if err != nil {
		// return error
		return err
	}

	// posts list to write
	posts := map[string]interface{}{}

	// loop through posts
	for rows.Next() {
		// scan
		var id, author int
		var created int64
		var content string
		err = rows.Scan(&id, &author, &created, &content)
		if err != nil {
			// return error
			return err
		}

		// post map to append
		post := map[string]interface{}{}
		post["id"] = id
		post["author"] = author
		post["created"] = created
		post["content"] = content

		// append post to posts map
		posts[strconv.Itoa(id)] = post
	}

	// write map
	return posts
}
