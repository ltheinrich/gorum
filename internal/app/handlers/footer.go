package handlers

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var (
	footerPage string
)

// Footer provide HTML page
func Footer(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check whether to load or not
	if footerPage == "" {
		// read footer page
		var footerData []byte
		footerData, err = ioutil.ReadFile("footer.html")

		// check for error
		if err != nil && !os.IsNotExist(err) {
			// print unknown error
			log.Println(err)
			return err
		}

		// set footer page variable
		footerPage = string(footerData)
	}

	// serve loaded footer
	if footerPage != "" {
		return map[string]interface{}{"footer": footerPage}
	}

	// return not footer found
	return errors.New("404 data")
}
