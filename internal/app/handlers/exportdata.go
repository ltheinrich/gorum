package handlers

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/NathanNr/gorum/internal/pkg/config"

	"github.com/NathanNr/gorum/internal/pkg/db"
)

// ExportData handler
func ExportData(data HandlerData) interface{} {
	var err error

	// check if authenticated
	if !data.Authenticated {
		// not authenticated
		return errors.New("403")
	}

	// get user table
	var userTable map[string]interface{}
	userTable, err = exportUserTable(data.Username)
	if err != nil {
		// return error
		fmt.Println("user")
		return err
	}

	// get tokens table
	var tokensTable map[int]map[string]interface{}
	tokensTable, err = exportTokensTable(data.Username)
	if err != nil {
		// return error
		fmt.Println("tokens")
		return err
	}

	// get userdata table
	var userDataTable map[int]map[string]interface{}
	userDataTable, err = exportUserDataTable(data.Username)
	if err != nil {
		// return error
		fmt.Println("userdata")
		return err
	}

	// get threads table
	var threadsTable map[int]map[string]interface{}
	threadsTable, err = exportThreadsTable(data.Username)
	if err != nil {
		// return error
		fmt.Println("threads")
		return err
	}

	// get posts table
	var postsTable map[int]map[string]interface{}
	postsTable, err = exportPostsTable(data.Username)
	if err != nil {
		// return error
		fmt.Println("posts")
		return err
	}

	return map[string]interface{}{"user": userTable, "tokens": tokensTable, "userdata": userDataTable, "threads": threadsTable, "posts": postsTable}
}

// export table data
func exportUserTable(username string) (table map[string]interface{}, err error) {
	// define variables
	var id int
	var passwordhash, registered, avatar string

	// query table
	err = db.DB.QueryRow(`SELECT id, username, passwordhash, registered FROM users WHERE username = $1;`,
		username).Scan(&id, &username, &passwordhash, &registered)
	if err != nil {
		return nil, err
	}

	// open avatar file
	var file *os.File
	file, err = os.Open(fmt.Sprintf("%v/%v.png", config.Get("data", "avatar"), id))
	defer file.Close()

	// check for error and if avatar file exists
	if err != nil && !os.IsNotExist(err) {
		log.Println(os.IsExist(err), os.IsNotExist(err))
		return nil, err
	} else if !os.IsNotExist(err) {
		// get avatar file info
		var fileInfo os.FileInfo
		fileInfo, err = file.Stat()
		if err != nil {
			return nil, err
		}

		// read avatar file
		avatarData := make([]byte, fileInfo.Size())
		_, err = io.ReadAtLeast(file, avatarData, int(fileInfo.Size()))
		if err != nil {
			return nil, err
		}

		// encode avatar to base64
		avatar = base64.StdEncoding.EncodeToString(avatarData)
	}

	// return data
	return map[string]interface{}{"id": id, "username": username, "passwordhash": passwordhash, "registered": registered, "avatar": avatar}, nil
}

// export table data
func exportTokensTable(username string) (table map[int]map[string]interface{}, err error) {
	// query table
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT tokens.token, tokens.holder, tokens.created FROM tokens
							INNER JOIN users ON tokens.holder = users.id WHERE users.username = $1;`,
		username)
	if err != nil {
		return nil, err
	}

	// loop through rows
	table = map[int]map[string]interface{}{}
	for i := 0; rows.Next(); i++ {
		// define variables
		var holder int
		var created int64
		var token string

		// read data
		err = rows.Scan(&token, &holder, &created)
		if err != nil {
			return nil, err
		}

		// set data in table
		table[i] = map[string]interface{}{}
		table[i]["token"] = token
		table[i]["holder"] = holder
		table[i]["created"] = created
	}

	// return data
	return table, nil
}

// export table data
func exportUserDataTable(username string) (table map[int]map[string]interface{}, err error) {
	// query table
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT userdata.holder, userdata.dataname, userdata.datavalue FROM userdata
							INNER JOIN users ON userdata.holder = users.id WHERE users.username = $1;`,
		username)
	if err != nil {
		return nil, err
	}

	// loop through rows
	table = map[int]map[string]interface{}{}
	for i := 0; rows.Next(); i++ {
		// define variables
		var holder int
		var dataname, datavalue string

		// read data
		err = rows.Scan(&holder, &dataname, &datavalue)
		if err != nil {
			return nil, err
		}

		// set data in table
		table[i] = map[string]interface{}{}
		table[i]["holder"] = holder
		table[i]["dataname"] = dataname
		table[i]["datavalue"] = datavalue
	}

	// return data
	return table, nil
}

// export table data
func exportThreadsTable(username string) (table map[int]map[string]interface{}, err error) {
	// query table
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT threads.id, threads.threadname, threads.board,
							threads.author, threads.created, threads.content FROM threads
							INNER JOIN users ON threads.author = users.id WHERE users.username = $1;`,
		username)
	if err != nil {
		return nil, err
	}

	// loop through rows
	table = map[int]map[string]interface{}{}
	for i := 0; rows.Next(); i++ {
		// define variables
		var id, board, author int
		var created int64
		var name, content string

		// read data
		err = rows.Scan(&id, &name, &board, &author, &created, &content)
		if err != nil {
			return nil, err
		}

		// set data in table
		table[i] = map[string]interface{}{}
		table[i]["id"] = id
		table[i]["name"] = name
		table[i]["board"] = board
		table[i]["author"] = author
		table[i]["created"] = created
		table[i]["content"] = content
	}

	// return data
	return table, nil
}

// export table data
func exportPostsTable(username string) (table map[int]map[string]interface{}, err error) {
	// query table
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT posts.id, posts.thread, posts.author, posts.created, posts.content FROM posts
							INNER JOIN users ON posts.author = users.id WHERE users.username = $1;`,
		username)
	if err != nil {
		return nil, err
	}

	// loop through rows
	table = map[int]map[string]interface{}{}
	for i := 0; rows.Next(); i++ {
		// define variables
		var id, thread, author int
		var created int64
		var content string

		// read data
		err = rows.Scan(&id, &thread, &author, &created, &content)
		if err != nil {
			return nil, err
		}

		// set data in table
		table[i] = map[string]interface{}{}
		table[i]["id"] = id
		table[i]["thread"] = thread
		table[i]["author"] = author
		table[i]["created"] = created
		table[i]["content"] = content
	}

	// return data
	return table, nil
}
