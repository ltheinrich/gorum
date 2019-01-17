package handlers

import (
	"errors"
	"log"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// DeletePost handler
func DeletePost(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check login
	if !auth {
		// not logged in
		return errors.New("403")
	}

	// get post id
	postID := GetInt(request, "postID")

	// check if data is provided
	if postID == 0 {
		// return not provided
		return errors.New("400")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM posts USING users WHERE posts.author = users.id
						AND posts.id = $1 AND users.username = $2;`, postID, username)
	if err != nil {
		log.Println(err)
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
