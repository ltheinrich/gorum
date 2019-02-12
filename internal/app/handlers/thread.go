package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ltheinrich/gorum/internal/pkg/config"
	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// Thread handler
func Thread(data HandlerData) interface{} {
	var err error

	// check if thread id provided
	threadID := data.Request.GetInt("threadID")
	if threadID == 0 {
		// thread id not provided
		return errors.New("400")
	}

	// define variables
	var created int64
	var board, author int
	var name, content, authorName string

	// query thread
	err = db.DB.QueryRow(`SELECT threads.threadname, threads.board, threads.author, threads.created, threads.content,
						users.username FROM threads INNER JOIN users ON threads.author = users.id WHERE threads.id = $1;`, threadID).
		Scan(&name, &board, &author, &created, &content, &authorName)

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// print and return error
		log.Println(err)
		return err
	}

	// thread map to write
	thread := map[string]interface{}{}
	thread["id"] = threadID
	thread["name"] = name
	thread["board"] = board
	thread["author"] = author
	thread["created"] = created
	thread["content"] = content
	thread["authorName"] = authorName

	// add avatar
	avatarPath := fmt.Sprintf("%s/%v.png", config.Get("data", "avatar"), author)
	_, err = os.Open(avatarPath)
	if os.IsNotExist(err) {
		thread["authorAvatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
	} else {
		thread["authorAvatar"] = avatarPath
	}

	// write map
	return thread
}
