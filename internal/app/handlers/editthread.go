package handlers

import (
	"errors"

	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// EditThread handler
func EditThread(data HandlerData) interface{} {
	var err error

	// check login
	if !data.Authenticated {
		// not logged in
		return errors.New("403")
	}

	// get variables from request
	threadID := data.Request.GetInt("threadID")
	title := data.Request.GetString("title")
	/* board := data.Request.GetInt("board") // TODO */
	content := data.Request.GetString("content")

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
	_, err = db.DB.Exec(`UPDATE threads SET threads.threadname = $1, threads.content = $2 FROM users
						WHERE threads.author = users.id AND threads.id = $3 AND users.username = $4;`,
		title, content, threadID, data.Username)
	if err != nil {
		// return error
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
