package handlers

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

var (
	footerPage string
)

// Footer provide HTML page
func Footer(data HandlerData) interface{} {
	var err error

	// check whether to load or not
	if footerPage == "" {
		// read footer page
		var footerData []byte
		footerData, err = ioutil.ReadFile("footer.html")

		// check for error
		if err != nil && !os.IsNotExist(err) {
			// return error
			return err
		}

		// set footer page variable
		footerPage = string(footerData)
	}

	// serve loaded footer
	if footerPage != "" {
		// query counts
		var users, threads, posts int
		db.DB.QueryRow(`SELECT (SELECT COUNT(*) FROM users),
						(SELECT COUNT(*) FROM threads),
						(SELECT COUNT(*) FROM posts);`).
			Scan(&users, &threads, &posts)

		// insert placeholders
		tempFooter := footerPage
		tempFooter = strings.Replace(tempFooter, "${users}", strconv.Itoa(users), -1)
		tempFooter = strings.Replace(tempFooter, "${threads}", strconv.Itoa(threads), -1)
		tempFooter = strings.Replace(tempFooter, "${posts}", strconv.Itoa(posts), -1)

		// return map
		return map[string]interface{}{"footer": tempFooter}
	}

	// return not footer found
	return errors.New("404 data")
}
