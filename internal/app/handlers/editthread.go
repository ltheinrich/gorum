package handlers

import (
	"errors"
	"log"

	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// EditThread handler
func EditThread(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check login
	if !auth {
		// not logged in
		return errors.New("403")
	}

	// get variables from request
	threadID := GetInt(request, "threadID")
	title := GetString(request, "title")
	/* board := GetInt(request, "board") // TODO */
	content := GetString(request, "content")

	// check if data is provided
	if threadID == 0 || title == "" || /* board == 0 || */ content == "" || len(title) > 32 {
		// return not provided
		return errors.New("400")
	}

	// check limit
	if len(content) > config.GetInt("limit", "thread") {
		// return too long
		return errors.New("411")
	}

	// insert into database
	_, err = db.DB.Exec(`UPDATE threads SET threadname = $1, content = $2 FROM users
						WHERE threads.id = $3 AND users.username = $4;`,
		title, content, threadID, username)
	if err != nil {
		// return unknown error
		log.Println(err)
		return err
	}

	// respond with id
	return map[string]interface{}{"done": true}
}
