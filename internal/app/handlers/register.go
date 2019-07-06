package handlers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dchest/captcha"

	"github.com/nathannr/gorum/internal/pkg/config"
	"github.com/nathannr/gorum/internal/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// Register handler
func Register(data HandlerData) interface{} {
	var err error

	// get strings from request
	password := data.Request.GetString("password")
	cap := data.Request.GetString("captcha")
	capVal := data.Request.GetString("captchaValue")

	// check if username and password are provided
	if data.Username == "" || password == "" || len(data.Username) > 32 {
		// return not provided
		return errors.New("400")
	}

	// verify captcha
	if config.Get("https", "captcha") == TRUE && !captcha.VerifyString(cap, capVal) {
		// invalid captcha
		return errors.New("403 captcha")
	}

	// query db
	var id int
	err = db.DB.QueryRow("SELECT id FROM users WHERE username = $1;", data.Username).Scan(&id)
	if err == sql.ErrNoRows {
		// not exists
		var passwordHash []byte
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+1)
		if err != nil {
			// return error
			return err
		}

		// insert into database
		_, err = db.DB.Exec("INSERT INTO users (username, passwordhash, registered) VALUES ($1, $2, $3);",
			data.Username, string(passwordHash), time.Now().Format("2006-01-02T15:04:05"))
		if err != nil {
			// return error
			return err
		}

		// generate new token
		token := newToken(data.Username)

		// registered
		return map[string]interface{}{"done": true, "token": token}
	} else if err != nil {
		// return error
		return err
	}

	// username exists
	return map[string]interface{}{"done": false}
}
