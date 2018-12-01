package handlers

import (
	"compress/gzip"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Data serve data files
func Data(rw http.ResponseWriter, r *http.Request) {
	var err error

	// check for malicious path
	path := strings.Replace(r.URL.Path, "/", "", 1)
	if strings.Contains(path, "..") {
		rw.WriteHeader(400)
		rw.Write([]byte{})
		return
	}

	// open file
	var file *os.File
	file, err = os.Open(path)
	defer file.Close()

	// get file extension
	extension := mime.TypeByExtension(filepath.Ext(path))
	if path == "/data/avatar/default" {
		extension = "image/png"
	}

	// set content-type and content-encoding
	rw.Header().Set("Content-Type", extension+"; charset=utf-8")
	rw.Header().Set("Content-Encoding", "gzip")

	// security headers
	SecurityHeaders(rw, r)

	// gzip compression
	w, _ := gzip.NewWriterLevel(rw, 2)
	defer w.Close()

	// check for error
	if err != nil {
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
