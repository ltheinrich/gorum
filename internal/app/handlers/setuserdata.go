package handlers

import (
	"errors"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// SetUserData handler
func SetUserData(data HandlerData) interface{} {
	var err error

	// check if authenticated
	if !data.Authenticated {
		// not authenticated
		return errors.New("403")
	}

	// check if data provided
	dataName := data.Request.GetString("dataName")
	dataValue := data.Request.GetString("dataValue")
	if dataName == "" || dataValue == "" {
		// both not provided
		return errors.New("400")
	}

	// update database
	_, err = db.DB.Exec(`INSERT INTO userdata (dataname, datavalue, holder)
						SELECT $1 AS dataname, $2 AS datavalue, id AS holder FROM users
						WHERE users.username = $3
						ON CONFLICT (holder, dataname) DO
						UPDATE SET datavalue = $2
						WHERE userdata.dataname = $1 AND userdata.holder = $4;`,
		dataName, dataValue, data.Username, GetUserID(data.Username))

	// check for error
	if err != nil {
		// return error
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
