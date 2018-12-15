package handlers

import (
	"database/sql"
	"errors"

	"github.com/dchest/captcha"
	"github.com/lheinrichde/gorum/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// Login handler
func Login(request map[string]interface{}, username string, auth bool) interface{} {
	// get strings from request
	cap := GetString(request, "captcha")
	capVal := GetString(request, "captchaValue")

	// verify captcha
	if !captcha.VerifyString(cap, capVal) {
		// invalid captcha
		return errors.New("403 captcha")
	}

	// write
	return map[string]interface{}{"valid": auth}
}

// verify login
func login(username, password string) bool {
	var err error

	// query db
	var passwordHash string
	err = db.DB.QueryRow("SELECT passwordhash FROM users WHERE username = $1 OR mail = $1;", username).Scan(&passwordHash)
	if err == sql.ErrNoRows {
		// not exists, but due to security return invalid
		return false
	}

	// compare passwords and return
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
