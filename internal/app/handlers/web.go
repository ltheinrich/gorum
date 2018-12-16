package handlers

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ltheinrich/gorum/pkg/config"
)

var (
	webFiles = map[string][]byte{}
)

// Web serve web/dist/gorum files
func Web(rw http.ResponseWriter, r *http.Request) {
	var err error

	// check for malicious path
	path := strings.Replace(r.URL.Path, "/", "", 1)
	if strings.Contains(path, "..") {
		rw.WriteHeader(400)
		rw.Write([]byte{})
		return
	}

	// define variables to open file
	var file *os.File
	filePath := fmt.Sprintf("%s/%s", config.Get("https", "directory"), path)

	// open file
	if path == "" {
		filePath += "index.html"
	}
	file, err = os.Open(filePath)

	// defer file close and set content-type and content-encoding
	defer file.Close()
	rw.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path))+"; charset=utf-8")
	rw.Header().Set("Content-Encoding", "gzip")

	// security headers
	SecurityHeaders(rw, r)

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
