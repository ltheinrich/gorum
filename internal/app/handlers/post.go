package handlers

import (
	"database/sql"
	"errors"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// Post handler
func Post(data HandlerData) interface{} {
	var err error

	// check if post id provided
	postID := data.Request.GetInt("postID")
	if postID == 0 {
		// post id not provided
		return errors.New("400")
	}

	// define variables
	var created int64
	var thread, author int
	var content, authorName string

	// query thread
	err = db.DB.QueryRow(`SELECT posts.thread, posts.author, posts.created, posts.content, users.username
						FROM posts INNER JOIN users ON posts.author = users.id WHERE posts.id = $1;`, postID).
		Scan(&thread, &author, &created, &content, &authorName)

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// return error
		return err
	}

	// post map to write
	post := map[string]interface{}{}
	post["id"] = postID
	post["thread"] = thread
	post["author"] = author
	post["created"] = created
	post["content"] = content
	post["authorName"] = authorName

	// write map
	return post
}
