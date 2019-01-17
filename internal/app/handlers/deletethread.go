package handlers

import (
	"errors"
	"log"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// DeleteThread handler
func DeleteThread(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check login
	if !auth {
		// not logged in
		return errors.New("403")
	}

	// get thread id
	threadID := GetInt(request, "threadID")

	// check if data is provided
	if threadID == 0 {
		// return not provided
		return errors.New("400")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM threads USING users WHERE threads.author = users.id
						AND threads.id = $1 AND users.username = $2;`, threadID, username)
	if err != nil {
		log.Println(err)
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
