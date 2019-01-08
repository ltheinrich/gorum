package handlers

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/ltheinrich/gorum/pkg/config"
	"github.com/ltheinrich/gorum/pkg/db"
)

// LastThreads handler
func LastThreads(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// get limit
	limit := GetInt(request, "limit")
	if limit == 0 || limit >= 20 {
		// no limit provided
		limit = 10
	}

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT threads.id, threads.threadname, threads.author, threads.created, users.username, posts.created
							FROM threads INNER JOIN users ON threads.author = users.id
							LEFT JOIN posts ON threads.id = posts.thread ORDER BY posts.created DESC;`)
	if err != nil {
		// return error
		return err
	}
	defer rows.Close()

	// threads list to write
	threads := map[string]interface{}{}

	// loop through threads
	for rows.Next() {
		// scan
		var id, author int
		var created int64
		var name, authorName string
		var answer interface{}
		err = rows.Scan(&id, &name, &author, &created, &authorName, &answer)
		if err != nil {
			// return error
			return err
		}
		idString := strconv.Itoa(id)

		// check if already exists
		if _, ok := threads[idString]; ok {
			continue
		} else if len(threads) >= limit {
			break
		}

		// thread map to append
		thread := map[string]interface{}{}
		thread["id"] = id
		thread["name"] = name
		thread["created"] = created
		thread["author"] = author
		thread["authorName"] = authorName

		// check if post exists
		if val, ok := answer.(int64); ok {
			thread["answer"] = val
		} else {
			thread["answer"] = created
		}

		// add avatar
		avatarPath := fmt.Sprintf("%s/%v.png", config.Get("data", "avatar"), author)
		_, err = os.Open(avatarPath)
		if os.IsNotExist(err) {
			thread["authorAvatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
		} else {
			thread["authorAvatar"] = avatarPath
		}
		fmt.Println(thread)
		// append thread to threads map
		threads[idString] = thread
	}

	// write map
	return threads
}
