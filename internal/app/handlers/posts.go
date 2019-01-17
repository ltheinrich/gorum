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

// Posts handler
func Posts(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// get thread id and check if provided
	threadID := GetInt(request, "threadID")
	if threadID == 0 {
		// no thread id provided
		return errors.New("400")
	}

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT posts.id, posts.author, posts.created, posts.content, users.username FROM posts
							INNER JOIN users ON posts.author = users.id WHERE posts.thread = $1;`, threadID)
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
		var content, authorName string
		err = rows.Scan(&id, &author, &created, &content, &authorName)
		if err != nil {
			// return error
			return err
		}

		// post map to append
		post := map[string]interface{}{}
		post["id"] = id
		post["author"] = author
		post["authorName"] = authorName
		post["created"] = created
		post["content"] = content

		// add avatar
		avatarPath := fmt.Sprintf("%s/%v.png", config.Get("data", "avatar"), author)
		_, err = os.Open(avatarPath)
		if os.IsNotExist(err) {
			post["authorAvatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
		} else {
			post["authorAvatar"] = avatarPath
		}

		// append post to posts map
		posts[strconv.Itoa(id)] = post
	}

	// write map
	return posts
}
