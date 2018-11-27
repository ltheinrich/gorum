package handlers

import (
	"encoding/base64"
	"mime/multipart"
	"net/http"

	"github.com/lheinrichde/gorum/pkg/db"
)

// UploadAvatar http handler function
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	var err error

	// security headers
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; img-src 'self' data:; style-src 'self' 'unsafe-inline';")

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

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

			// encode and update db
			image := base64.StdEncoding.EncodeToString(fileData)
			_, err = db.DB.Exec("UPDATE users SET avatar = $1 WHERE username = $2;", image, username)
			if err != nil {
				// write error
				w.Write([]byte(err.Error()))
			}
			http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
		}
	}
}
