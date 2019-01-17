package handlers

import (
	"errors"
	"time"

	"github.com/ltheinrich/captcha"
	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// NewPost handler
func NewPost(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check login
	if !auth {
		// not logged in
		return errors.New("403")
	}

	// get strings from request
	thread := GetInt(request, "thread")
	content := GetString(request, "content")
	cap := GetString(request, "captcha")
	capVal := GetString(request, "captchaValue")

	// check if data is provided
	if thread == 0 || content == "" {
		// return not provided
		return errors.New("400")
	}

	// verify captcha
	if config.Get("https", "captcha") == TRUE && !captcha.VerifyString(cap, capVal) {
		// invalid captcha
		return errors.New("403 captcha")
	}

	// insert into database
	_, err = db.DB.Exec("INSERT INTO posts (thread, author, created, content) VALUES ($1, $2, $3, $4);",
		thread, GetUserID(username), time.Now().Unix(), content)
	if err != nil {
		// return error
		return err
	}

	// respond done
	return map[string]interface{}{"done": true}
}
