package handlers

import (
	"database/sql"
	"log"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// Boards handler
func Boards(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT boards.id, boards.boardname, boards.boarddescription,
							boards.boardicon, categories.categoryname,
							boards.sort + categories.sort FROM boards
							INNER JOIN categories ON boards.category = categories.id;`)

	// defer close and check for error
	defer rows.Close()
	if err != nil {
		// print and return error
		log.Println(err)
		return err
	}

	// boards list to write
	categories := map[string]interface{}{}

	// loop through boards
	for rows.Next() {
		// scan
		var id, sort int
		var name, description, icon, category string
		err = rows.Scan(&id, &name, &description, &icon, &category, &sort)
		if err != nil {
			// print and return error
			log.Println(err)
			return err
		}

		// board map to append
		board := map[string]interface{}{}
		board["id"] = id
		board["name"] = name
		board["description"] = description
		board["icon"] = icon
		board["sort"] = sort

		// append board to categories map
		if categories[category] == nil {
			categories[category] = []map[string]interface{}{}
		}
		categories[category] = append(categories[category].([]map[string]interface{}), board)
	}

	// write map
	return categories
}
