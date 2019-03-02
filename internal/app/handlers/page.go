package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ltheinrich/gorum/internal/pkg/config"
)

var (
	pages = map[string]string{}
)

// Page provide custom pages
func Page(data HandlerData) interface{} {
	var err error

	// get page name
	rawName := strings.ToLower(data.Request.GetString("name"))
	_, name := filepath.Split(rawName)
	if name == "" {
		// name not provided
		return errors.New("400")
	}

	// check whether to load or not
	if _, exists := pages[name]; !exists {
		// read page
		var data []byte
		pagePath := fmt.Sprintf("%v/%v.html", config.Get("data", "page"), name)
		data, err = ioutil.ReadFile(pagePath)

		// check for error
		if err != nil && !os.IsNotExist(err) {
			// return error
			return err
		}

		// set pages map variable
		pages[name] = string(data)
	}

	// serve loaded page
	if len(pages[name]) > 0 {
		// return page
		return map[string]interface{}{"page": pages[name]}
	}

	// return page not found
	return errors.New("404 data")
}
