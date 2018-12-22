package handlers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ltheinrich/captcha"
	"github.com/ltheinrich/gorum/pkg/config"
	"github.com/ltheinrich/gorum/pkg/db"
)

// NewThread handler
func NewThread(request map[string]interface{}, username string, auth bool) interface{} {
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

	// verify captcha
	if config.Get("https", "captcha") == "true" && !captcha.VerifyString(cap, capVal) {
		// invalid captcha
		return errors.New("403 captcha")
	}

	// insert into database
	var id int
	var row *sql.Row
	row = db.DB.QueryRow("INSERT INTO threads (threadname, board, author, created, content) VALUES ($1, $2, $3, $4, $5) RETURNING id;",
		title, board, GetUserID(username), time.Now().Unix(), content)
	row.Scan(&id)

	// respond with id
	return map[string]interface{}{"id": id}
}
