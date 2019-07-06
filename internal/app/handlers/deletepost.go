package handlers

import (
	"errors"

	"github.com/NathanNr/gorum/internal/pkg/db"
)

// DeletePost handler
func DeletePost(data HandlerData) interface{} {
	var err error

	// check login
	if !data.Authenticated {
		// not logged in
		return errors.New("403")
	}

	// get post id
	postID := data.Request.GetInt("postID")

	// check if data is provided
	if postID == 0 {
		// return not provided
		return errors.New("400")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM posts USING users WHERE posts.author = users.id
						AND posts.id = $1 AND users.username = $2;`, postID, data.Username)
	if err != nil {
		// return error
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
