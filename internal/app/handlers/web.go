package handlers

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ltheinrich/gorum/internal/pkg/webassets"
)

var (
	webFiles         = map[string][]byte{}
	customFavicon    []byte
	customFaviconPNG []byte
	customTouchIcon  []byte
)

// Web serve web/dist/gorum files
func Web(rw http.ResponseWriter, r *http.Request) {
	var err error

	// data to respond with
	var file []byte

	// deliver custom images
	path := strings.Replace(r.URL.Path, "/", "", 1)
	file = customImages(path)

	// set content-type and content-encoding
	rw.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path))+"; charset=utf-8")
	rw.Header().Set("Content-Encoding", "gzip")

	// security headers
	SecurityHeaders(rw, r)

	// gzip compression
	w, _ := gzip.NewWriterLevel(rw, 2)
	defer w.Close()

	if file == nil {
		// get file
		file, err = webassets.Asset(path)

		// if not exists load index.html
		if err != nil {
			file = webassets.MustAsset("index.html")
		}
	}

	// write file
	_, err = w.Write(file)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
}

// custom images to deliver
func customImages(path string) (file []byte) {
	if path == "favicon.ico" {
		// check if exists
		if customFavicon == nil {
			// read file
			file, err := ioutil.ReadFile(path)
			if err == nil {
				// set custom icon and return file
				customFavicon = file
				return customFavicon
			}

			// set not exists and return null
			customFavicon = []byte{0}
			return nil
		} else if len(customFavicon) > 2 {
			// already exists, return
			return customFavicon
		}
	} else if path == "apple-touch-icon.png" {
		// check if exists
		if customFaviconPNG == nil {
			// read file
			file, err := ioutil.ReadFile(path)
			if err == nil {
				// set custom icon and return file
				customFaviconPNG = file
				return customFaviconPNG
			}
			// set not exists and return null
			customFaviconPNG = []byte{0}
			return nil
		} else if len(customFaviconPNG) > 2 {
			// already exists, return
			return customFaviconPNG
		}
	} else if path == "favicon.png" {
		// check if exists
		if customTouchIcon == nil {
			// read file
			file, err := ioutil.ReadFile(path)
			if err == nil {
				// set custom icon and return file
				customTouchIcon = file
				return customTouchIcon
			}

			// set not exists and return null
			customTouchIcon = []byte{0}
			return nil
		} else if len(customTouchIcon) > 2 {
			// already exists, return
			return customTouchIcon
		}
	}

	return nil
}
