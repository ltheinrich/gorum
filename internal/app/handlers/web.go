package handlers

import (
	"compress/gzip"
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
func Web(rw http.ResponseWriter, r *http.Request) {
	var err error

	// get filename
	_, fileName := path.Split(r.URL.Path)
	if fileName == "" || fileName == "/" {
		fileName = "index.html"
	}

	// open file
	var file *os.File
	file, err = os.Open(fmt.Sprintf("%s/%s", config.Get("https", "directory"), fileName))

	// defer file close and set content-type and content-encoding
	defer file.Close()
	rw.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(fileName))+"; charset=utf-8")
	rw.Header().Set("Content-Encoding", "gzip")

	// gzip compression
	w, _ := gzip.NewWriterLevel(rw, 2)
	defer w.Close()

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

	// write file
	_, err = io.Copy(w, file)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
}
