package handlers

import (
	"errors"

	"github.com/nathannr/gorum/internal/pkg/db"
)

// DeleteThread handler
func DeleteThread(data HandlerData) interface{} {
	var err error

	// check login
	if !data.Authenticated {
		// not logged in
		return errors.New("403")
	}

	// get thread id
	threadID := data.Request.GetInt("threadID")

	// check if data is provided
	if threadID == 0 {
		// return not provided
		return errors.New("400")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM threads USING users WHERE threads.author = users.id
						AND threads.id = $1 AND users.username = $2;`, threadID, data.Username)
	if err != nil {
		// return error
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
