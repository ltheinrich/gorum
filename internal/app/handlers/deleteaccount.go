package handlers

import (
	"errors"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// DeleteAccount handler
func DeleteAccount(data HandlerData) interface{} {
	var err error

	// authenticate
	password := data.Request.GetString("password")
	if password == "" || !data.Authenticated || !login(data.Username, password) {
		// not authenticated
		return errors.New("403")
	}

	// delete account
	_, err = db.DB.Exec("DELETE FROM users WHERE username = $1;", data.Username)
	if err != nil {
		// return error
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
