package handlers

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/NathanNr/gorum/internal/pkg/config"
	"github.com/NathanNr/gorum/internal/pkg/db"
	"github.com/nfnt/resize"
)

// UploadAvatar http handler function
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	var err error
	defer r.Body.Close()

	// security headers
	SecurityHeaders(w, r)

	// get username and password
	rawUsername, _ := base64.URLEncoding.DecodeString(r.FormValue("username"))
	rawToken, _ := base64.URLEncoding.DecodeString(r.FormValue("token"))
	username, token := string(rawUsername), string(rawToken)

	// check if provided and login is correct
	if username != "" && token != "" && validateToken(username, token) {
		// get file
		var file multipart.File
		var header *multipart.FileHeader
		file, header, err = r.FormFile("avatar")

		// defer close file
		defer file.Close()

		// check if file provided
		if err != nil || file == nil || header == nil {
			// write header
			w.Header().Add("content-type", "text/html")
			w.WriteHeader(200)

			// write content
			w.Write([]byte(`<form method="post" enctype="multipart/form-data"><input name="avatar" type="file" size="50" accept=".png,.jpg,.jpeg"><button type="submit">Avatar Upload</button></form>`))
		} else {
			// check avatar size limit
			var avatarSizeLimit int
			avatarSizeLimit, _ = strconv.Atoi(config.Get("limit", "avatar"))
			if int(header.Size) > avatarSizeLimit {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// write content and return
				w.Write([]byte(`<h3>Avatar size limit exceeded</h3><form method="post" enctype="multipart/form-data"><input name="avatar" type="file" size="50" accept=".png,.jpg,.jpeg"><button type="submit">Avatar Upload</button></form>`))
				return
			}

			/*
				// read file
				fileData := make([]byte, header.Size)
				io.ReadAtLeast(file, fileData, int(header.Size))
			*/

			// decode avatar image
			var img image.Image
			if strings.HasSuffix(header.Filename, ".png") {
				img, err = png.Decode(file)
			} else if strings.HasSuffix(header.Filename, ".jpg") || strings.HasSuffix(header.Filename, ".jpeg") {
				img, err = jpeg.Decode(file)
			} else {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// print and write error
				w.Write([]byte(`<h3>Wrong file extension, use .png, .jpg or .jpeg</h3><form method="post" enctype="multipart/form-data"><input name="avatar" type="file" size="50" accept=".png,.jpg,.jpeg"><button type="submit">Avatar Upload</button></form>`))
				return
			}

			// check for error
			if err != nil {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// print and write error
				log.Println(err)
				w.Write([]byte(err.Error()))
				return
			}

			// get user id
			var userID int
			err = db.DB.QueryRow("SELECT id from users WHERE username = $1;", username).Scan(&userID)
			if err != nil {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// print and write error
				log.Println(err)
				w.Write([]byte(err.Error()))
				return
			}

			// open avatar file
			var avatar *os.File
			avatarName := fmt.Sprintf("%v/%v.png", config.Get("data", "avatar"), userID)
			avatar, err = os.OpenFile(avatarName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)

			// defer close avatar file
			defer avatar.Close()

			// create directories
			if os.IsNotExist(err) {
				err = os.MkdirAll(config.Get("data", "avatar"), os.ModePerm)
				if err != nil {
					// write header
					w.Header().Add("content-type", "text/html")
					w.WriteHeader(200)

					// print and write error
					log.Println(err)
					w.Write([]byte(err.Error()))
					return
				}

				// open avatar file again and defer close
				avatar, err = os.OpenFile(avatarName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
				defer avatar.Close()

				// check for error
				if err != nil {
					// write header
					w.Header().Add("content-type", "text/html")
					w.WriteHeader(200)

					// print and write error
					log.Println(err)
					w.Write([]byte(err.Error()))
					return
				}
			} else if err != nil {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// print and write error
				log.Println(err)
				w.Write([]byte(err.Error()))
				return
			}

			// resize and write avatar file
			img = resize.Resize(100, 100, img, resize.Bilinear)
			err = png.Encode(avatar, img)
			if err != nil {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// print and write error
				log.Println(err)
				w.Write([]byte(err.Error()))
				return
			}

			// redirect
			http.Redirect(w, r, "/edit-profile", http.StatusSeeOther)
		}
	}
}
