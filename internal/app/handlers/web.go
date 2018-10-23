package handlers

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/lheinrichde/gorum/pkg/config"
)

var (
	webFiles = map[string][]byte{}
)

// Web serve web/dist/gorum files
func Web(w http.ResponseWriter, r *http.Request) {
	var err error

	// get filename
	_, fileName := path.Split(r.URL.Path)
	if fileName == "" || fileName == "/" {
		fileName = "index.html"
	}

	// open file
	var file *os.File
	file, err = os.Open(fmt.Sprintf("%s/%s", config.Get("https", "directory"), fileName))

	// check if file exists
	if os.IsNotExist(err) {
		// get index file
		file, err = os.Open(fmt.Sprintf("%s/index.html", config.Get("https", "directory")))
		if err != nil {
			// unknown error
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
	} else if err != nil {
		// unknown error
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	// defer file close and set content-type
	defer file.Close()
	w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(fileName))+"; charset=utf-8")

	// write file
	_, err = io.Copy(w, file)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
}
