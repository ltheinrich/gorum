package handlers

import (
	"database/sql"
	"errors"

	"github.com/ltheinrich/gorum/pkg/db"
)

// Thread handler
func Thread(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check if thread id provided
	threadID := GetInt(request, "threadID")
	if threadID == 0 {
		// thread id not provided
		return errors.New("400")
	}

	// define variables
	var created int64
	var board, author int
	var name, content string

	// query thread
	err = db.DB.QueryRow("SELECT threadname, board, author, created, content FROM threads WHERE id = $1;", threadID).
		Scan(&name, &board, &author, &created, &content)

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// return error
		return err
	}

	// thread map to write
	thread := map[string]interface{}{}
	thread["id"] = threadID
	thread["name"] = name
	thread["board"] = board
	thread["author"] = author
	thread["created"] = created
	thread["content"] = content

	// write map
	return thread
}
