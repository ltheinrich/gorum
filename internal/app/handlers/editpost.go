package handlers

import (
	"errors"

	"github.com/nathannr/gorum/internal/pkg/config"
	"github.com/nathannr/gorum/internal/pkg/db"
)

// EditPost handler
func EditPost(data HandlerData) interface{} {
	var err error

	// check login
	if !data.Authenticated {
		// not logged in
		return errors.New("403")
	}

	// get variables from request
	postID := data.Request.GetInt("postID")
	content := data.Request.GetString("content")

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
		content, postID, data.Username)
	if err != nil {
		// return error
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
