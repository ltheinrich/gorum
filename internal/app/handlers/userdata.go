package handlers

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/nathannr/gorum/internal/pkg/db"
)

// UserData handler
func UserData(data HandlerData) interface{} {
	var err error

	// check if data provided
	userID := data.Request.GetInt("userID")
	dataNames := data.Request.GetStringArray("dataNames")
	if len(dataNames) == 0 || (userID == 0 && data.Username == "") {
		// both not provided
		return errors.New("400")
	}

	// define query variables
	query := "userdata.dataname = $2"
	injections := []interface{}{dataNames[0]}

	// loop thorugh data names, start with the second
	for i := 1; i < len(dataNames); i++ {
		// extend query and add data name to injections
		query += " OR userdata.dataname = $" + strconv.Itoa(i+2)
		injections = append(injections, dataNames[i])
	}

	// rows variable
	var rows *sql.Rows

	// check what provided
	if userID == 0 {
		// query current user
		rows, err = db.DB.Query(`SELECT userdata.dataname, userdata.datavalue FROM userdata INNER JOIN users
							ON userdata.holder = users.id WHERE users.username = $1 AND (`+query+`);`,
			append([]interface{}{data.Username}, injections...)...)
	} else {
		// query user by id
		rows, err = db.DB.Query(`SELECT userdata.dataname, userdata.datavalue FROM userdata WHERE userdata.holder = $1 AND (`+query+`);`,
			append([]interface{}{userID}, injections...)...)
	}

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// return error
		return err
	}

	// define userdata map annd loop through rows
	userdata := map[string]interface{}{}
	for rows.Next() {
		// scan row
		var queryName, queryValue string
		err = rows.Scan(&queryName, &queryValue)

		// check not found
		if err == sql.ErrNoRows {
			// return not found
			return errors.New("404")
		} else if err != nil {
			// return error
			return err
		}

		// insert user data
		userdata[queryName] = queryValue
	}

	// write map
	return userdata
}
