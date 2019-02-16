package handlers

import (
	"database/sql"
	"errors"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// Board handler
func Board(data HandlerData) interface{} {
	var err error

	// check if board id provided
	boardID := data.Request.GetInt("boardID")
	if boardID == 0 {
		// not provided
		return errors.New("400")
	}

	// define variables
	var sort, category int
	var name, description, icon string

	// query board by id
	err = db.DB.QueryRow("SELECT boardname, boarddescription, boardicon, sort, category FROM boards WHERE id = $1;", boardID).
		Scan(&name, &description, &icon, &sort, &category)

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// return error
		return err
	}

	// board map to write
	board := map[string]interface{}{}
	board["id"] = boardID
	board["name"] = name
	board["description"] = description
	board["icon"] = icon
	board["sort"] = sort
	board["category"] = category

	// write map
	return board
}
