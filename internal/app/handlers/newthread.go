package handlers

import (
	"errors"
	"time"

	"github.com/ltheinrich/captcha"
	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// NewThread handler
func NewThread(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check login
	if !auth {
		// not logged in
		return errors.New("403")
	}

	// get strings from request
	title := GetString(request, "title")
	board := GetInt(request, "board")
	content := GetString(request, "content")
	cap := GetString(request, "captcha")
	capVal := GetString(request, "captchaValue")

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
		title, board, GetUserID(username), time.Now().Unix(), content).Scan(&id)
	if err != nil {
		// return error
		return err
	}

	// respond with id
	return map[string]interface{}{"id": id}
}
