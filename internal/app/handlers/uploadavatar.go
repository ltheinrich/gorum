package handlers

import (
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/lheinrichde/gorum/pkg/config"
)

// UploadAvatar http handler function
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	var err error

	// security headers
	SecurityHeaders(w, r)

	// get username and password
	rawUsername, _ := base64.StdEncoding.DecodeString(r.FormValue("username"))
	username, password := string(rawUsername), r.FormValue("password")

	// check if provided and login is correct
	if username != "" && password != "" && login(username, password) {
		// get file
		var file multipart.File
		var header *multipart.FileHeader
		file, header, err = r.FormFile("avatar")

		// check if file provided
		if err != nil || file == nil || header == nil {
			// write header
			w.Header().Add("content-type", "text/html")
			w.WriteHeader(200)

			// write content
			w.Write([]byte(`<form method="post" enctype="multipart/form-data"><input name="avatar" type="file" size="50" accept="image/*"><button type="submit">Avatar</button></form>`))
		} else {
			// read file
			fileData := make([]byte, header.Size)
			file.Read(fileData)
			defer file.Close()

			// open avatar file
			var avatar *os.File
			avatarName := fmt.Sprintf("%s/%s.png", config.Get("data", "avatar"), username)
			avatar, err = os.OpenFile(avatarName, os.O_RDWR|os.O_CREATE, os.ModePerm)

			// create directories
			if os.IsNotExist(err) {
				err = os.MkdirAll(config.Get("data", "avatar"), os.ModePerm)
				if err != nil {
					// write header
					w.Header().Add("content-type", "text/html")
					w.WriteHeader(200)

					// write content
					w.Write([]byte(err.Error()))
					return
				}

				// open avatar file again
				avatar, err = os.OpenFile(avatarName, os.O_RDWR|os.O_CREATE, os.ModePerm)
				if err != nil {
					// write header
					w.Header().Add("content-type", "text/html")
					w.WriteHeader(200)

					// write content
					w.Write([]byte(err.Error()))
					return
				}
			} else if err != nil {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// write content
				w.Write([]byte(err.Error()))
				return
			}

			// write avatar file
			_, err := avatar.Write(fileData)
			if err != nil {
				// write header
				w.Header().Add("content-type", "text/html")
				w.WriteHeader(200)

				// write content
				w.Write([]byte(err.Error()))
				return
			}

			// redirect
			http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
		}
	}
}
