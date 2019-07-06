package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/NathanNr/gorum/internal/pkg/config"
	"github.com/NathanNr/gorum/internal/pkg/db"
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
	var name, content, authorName, boardName string

	// query thread
	err = db.DB.QueryRow(`SELECT threads.threadname, threads.board, threads.author, threads.created, threads.content,
						users.username, boards.boardname FROM threads INNER JOIN users ON threads.author = users.id
						INNER JOIN boards ON boards.id = threads.board WHERE threads.id = $1;`, threadID).
		Scan(&name, &board, &author, &created, &content, &authorName, &boardName)

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// return error
		return err
	}

	// thread map to write
	thread := map[string]interface{}{}
	thread["id"] = threadID
	thread["name"] = name
	thread["board"] = board
	thread["boardName"] = boardName
	thread["author"] = author
	thread["created"] = created
	thread["content"] = content
	thread["authorName"] = authorName

	// add avatar
	avatarPath := fmt.Sprintf("%v/%v.png", config.Get("data", "avatar"), author)
	_, err = os.Open(avatarPath)
	if os.IsNotExist(err) {
		thread["authorAvatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
	} else {
		thread["authorAvatar"] = avatarPath
	}

	// write map
	return thread
}
