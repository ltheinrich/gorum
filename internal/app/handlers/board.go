package handlers

import (
	"database/sql"
	"errors"

	"github.com/lheinrichde/gorum/pkg/db"
)

// Board handler
func Board(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check if board id provided
	boardID := GetInt(request, "boardID")
	if boardID == 0 {
		// not provided
		return errors.New("400")
	}

	// define variables
	var name, board, author, created string

	// query user by id
	err = db.DB.QueryRow("SELECT threadname, board, author, created FROM boards WHERE id = $1;", boardID).Scan(&name, &board, &author, &created)

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// return error
		return err
	}

	// board map to write
	thread := map[string]interface{}{}
	thread["id"] = boardID
	thread["name"] = name
	thread["board"] = board
	thread["author"] = author
	thread["created"] = created

	// write map
	return thread
}
