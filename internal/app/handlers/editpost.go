package handlers

import (
	"errors"
	"log"

	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// EditPost handler
func EditPost(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check login
	if !auth {
		// not logged in
		return errors.New("403")
	}

	// get variables from request
	postID := GetInt(request, "postID")
	content := GetString(request, "content")

	// check if data is provided
	if postID == 0 || content == "" {
		// return not provided
		return errors.New("400")
	}

	// check limit
	if len(content) > config.GetInt("limit", "post") {
		// return too long
		return errors.New("411")
	}

	// insert into database
	_, err = db.DB.Exec(`UPDATE posts SET content = $1 FROM users
						WHERE posts.author = users.id AND posts.id = $2 AND users.username = $3;`,
		content, postID, username)
	if err != nil {
		// print and return error
		log.Println(err)
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
