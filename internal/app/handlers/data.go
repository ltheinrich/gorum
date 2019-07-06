package handlers

import (
	"compress/gzip"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/NathanNr/gorum/internal/pkg/assets"
)

// Data serve data files
func Data(rw http.ResponseWriter, r *http.Request) {
	var err error

	// get path without slash
	path := r.URL.Path[1:]

	// set content-type and content-encoding
	rw.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path))+"; charset=utf-8")
	rw.Header().Set("Content-Encoding", "gzip")

	// security headers
	SecurityHeaders(rw, r)

	// gzip compression
	w, _ := gzip.NewWriterLevel(rw, 2)
	defer w.Close()

	// default paths
	switch path {
	case "data/avatar/default":
		// write default avatar
		_, err = w.Write(assets.MustAsset("avatar.png"))

		// check for error
		if err != nil {
			// print unknown error
			log.Println(err)
		}

		// return
		return
	}

	// open file and defer close
	var file *os.File
	file, err = os.Open(path)
	defer file.Close()

	// check for error
	if err != nil {
		// print unknown error
		log.Println(err)
		return
	}

	// write file
	_, err = io.Copy(w, file)
	if err != nil {
		// print unknown error
		log.Println(err)
	}
}
