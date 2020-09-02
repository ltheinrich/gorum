package handlers

import (
	"errors"
	"time"

	"github.com/dchest/captcha"
	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// NewThread handler
func NewThread(data HandlerData) interface{} {
	var err error

	// check login
	if !data.Authenticated {
		// not logged in
		return errors.New("403")
	}

	// get strings from request
	title := data.Request.GetString("title")
	board := data.Request.GetInt("board")
	content := data.Request.GetString("content")
	cap := data.Request.GetString("captcha")
	capVal := data.Request.GetString("captchaValue")

	// check if data is provided
	if title == "" || board == 0 || content == "" || len(title) > 32 {
		// return not provided
		return errors.New("400")
	}

	// check limit
	if len(content) > config.GetInt("limit", "thread") {
		// return too long
		return errors.New("411")
	}

	// verify captcha
	if config.Get("https", "captcha") == TRUE && !captcha.VerifyString(cap, capVal) {
		// invalid captcha
		return errors.New("403 captcha")
	}

	// insert into database
	var id int
	err = db.DB.QueryRow(`INSERT INTO threads (threadname, board, author, created, content)
												VALUES ($1, $2, $3, $4, $5) RETURNING id;`,
		title, board, GetUserID(data.Username), time.Now().Unix(), content).Scan(&id)
	if err != nil {
		// return error
		return err
	}

	// respond with id
	return map[string]interface{}{"id": id}
}
