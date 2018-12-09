package handlers

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/lheinrichde/gorum/pkg/db"
)

// Threads handler
func Threads(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// get board id and check if provided
	boardID := GetInt(request, "boardID")
	if boardID == 0 {
		// no board id provided
		return errors.New("400")
	}

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT id, threadname, author, created FROM threads WHERE board = $1;`, boardID)
	if err != nil {
		// return error
		return err
	}

	// threads list to write
	threads := map[string]interface{}{}

	// loop through threads
	for rows.Next() {
		// scan
		var id, author int
		var created int64
		var name string
		err = rows.Scan(&id, &name, &author, &created)
		if err != nil {
			// return error
			return err
		}

		// thread map to append
		thread := map[string]interface{}{}
		thread["id"] = id
		thread["name"] = name
		thread["created"] = created
		thread["author"] = author

		// append thread to threads map
		threads[strconv.Itoa(id)] = thread
	}

	// write map
	return threads
}
