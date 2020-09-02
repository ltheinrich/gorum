package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/dchest/captcha"
	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// Login handler
func Login(data HandlerData) interface{} {
	// check if already authenticated using a token
	if data.Authenticated {
		// return valid
		return map[string]interface{}{"valid": data.Authenticated}
	}

	// get password and token string
	password := data.Request.GetString("password")
	token := data.Request.GetString("token")
	if password == "" && token == "" {
		// password missing
		return errors.New("400")
	} else if token != "" {
		// return invalid
		return map[string]interface{}{"valid": false}
	}

	// get captcha values
	cap := data.Request.GetString("captcha")
	capVal := data.Request.GetString("captchaValue")

	// verify captcha
	if config.Get("https", "captcha") == TRUE && !captcha.VerifyString(cap, capVal) {
		// invalid captcha
		return errors.New("403 captcha")
	}

	// validate password
	validPassword := login(data.Username, password)
	if !validPassword {
		// return wrong password
		return errors.New("403")
	}

	// generate and return new token
	return map[string]interface{}{"token": newToken(data.Username)}
}

// verify login using password
func login(username, password string) bool {
	var err error

	// query db
	var passwordHash string
	err = db.DB.QueryRow("SELECT passwordhash FROM users WHERE username = $1;", username).Scan(&passwordHash)
	if err == sql.ErrNoRows {
		// not exists, but due to security return invalid
		return false
	}

	// compare passwords and return
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

// validate token for username
func validateToken(username, token string) (valid bool) {
	var err error

	// query db and check for error
	var dbToken string
	err = db.DB.QueryRow(`SELECT tokens.token FROM tokens INNER JOIN users
					ON tokens.holder = users.id WHERE tokens.token = $1 AND users.username = $2;`, token, username).Scan(&dbToken)

	// check if token does not exist
	if err == sql.ErrNoRows {
		// non-existent
		return false
	} else if err != nil {
		log.Println(err)
		return false
	}

	// check if database token matches provided token
	return token == dbToken
}

// generate and store new token
func newToken(username string) (token string) {
	var err error

	// generate token bytes
	tokenData := make([]byte, 64)
	_, err = rand.Read(tokenData)
	if err != nil {
		log.Println(err)
		return ""
	}

	// encode token with base64
	token = base64.StdEncoding.EncodeToString(tokenData)

	// delete old tokens from database
	_, err = db.DB.Exec(`DELETE FROM tokens WHERE token NOT IN
						(SELECT token FROM tokens ORDER BY created DESC LIMIT $1);`,
		config.Get("limit", "tokens"))
	if err != nil {
		log.Println(err)
		return ""
	}

	// store token in database
	_, err = db.DB.Exec(`INSERT INTO tokens (token, holder, created)
						SELECT $1 AS token, users.id AS holder, $2 AS created FROM users WHERE users.username = $3;`,
		token, time.Now().Unix(), username)
	if err != nil {
		log.Println(err)
		return ""
	}

	// return token
	return token
}
