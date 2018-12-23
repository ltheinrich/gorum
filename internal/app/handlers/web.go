package handlers

import (
	"compress/gzip"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ltheinrich/gorum/internal/pkg/webassets"
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

	// set content-type and content-encoding
	rw.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path))+"; charset=utf-8")
	rw.Header().Set("Content-Encoding", "gzip")

	// security headers
	SecurityHeaders(rw, r)

	// gzip compression
	w, _ := gzip.NewWriterLevel(rw, 2)
	defer w.Close()

	// get file
	var file []byte
	file, err = webassets.Asset(path)

	// if not exists load index.html
	if err != nil {
		file = webassets.MustAsset("index.html")
	}

	// write file
	_, err = w.Write(file)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
}
