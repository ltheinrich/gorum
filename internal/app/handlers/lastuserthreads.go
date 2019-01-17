package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// LastUserThreads handler
func LastUserThreads(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// get limit
	limit := GetInt(request, "limit")
	if limit == 0 || limit >= 20 {
		// no limit provided
		limit = 10
	}

	// get user id and check if provided
	userID := GetInt(request, "userID")
	if userID == 0 {
		// not provided
		return errors.New("400")
	}

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT threads.id, threads.threadname, threads.author, threads.board, threads.created, users.username, posts.created
							FROM threads INNER JOIN users ON threads.author = users.id
							LEFT JOIN posts ON threads.id = posts.thread AND posts.author = $1
							WHERE posts.author = $1 OR threads.author = $1 ORDER BY posts.created DESC;`, userID)
	if err != nil {
		// return error
		return err
	}
	defer rows.Close()

	// threads list to write
	threads := map[string]interface{}{}

	// loop through threads
	for rows.Next() {
		// scan
		var id, board, author int
		var created int64
		var name, authorName string
		var answer interface{}
		err = rows.Scan(&id, &name, &author, &board, &created, &authorName, &answer)
		if err != nil {
			// return error
			return err
		}
		idString := strconv.Itoa(id)

		// check if already exists
		if _, ok := threads[idString]; ok {
			continue
		} else if len(threads) >= limit {
			break
		}

		// thread map to append
		thread := map[string]interface{}{}
		thread["id"] = id
		thread["name"] = name
		thread["created"] = created
		thread["board"] = board
		thread["authorName"] = authorName

		// get avatar path
		avatarPath := fmt.Sprintf("%s/%v.png", config.Get("data", "avatar"), author)
		_, err = os.Open(avatarPath)
		if os.IsNotExist(err) {
			thread["authorAvatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
		} else {
			thread["authorAvatar"] = avatarPath
		}

		// check if post exists
		if val, ok := answer.(int64); ok {
			thread["answer"] = val
		} else {
			thread["answer"] = created
		}

		// append thread to threads map
		threads[idString] = thread
	}

	// write map
	return threads
}
