package handlers

import (
	"net/http"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

const (
	// TRUE string ("true")
	TRUE = "true"

	// FALSE string ("false")
	FALSE = "false"
)

// GetString get string or empty string
func GetString(request map[string]interface{}, name string) string {
	// cast and return value
	value, _ := request[name].(string)
	return value
}

// GetStringArray get string array or nil
func GetStringArray(request map[string]interface{}, name string) []string {
	// cast and return value
	value, _ := request[name].([]string)
	return value
}

// GetInt get int or 0
func GetInt(request map[string]interface{}, name string) int {
	// cast and return value
	value, _ := request[name].(int)
	return value
}

// GetBool get string or false
func GetBool(request map[string]interface{}, name string) bool {
	// cast and return value
	value, _ := request[name].(bool)
	return value
}

// SecurityHeaders write CORS and CSP headers
func SecurityHeaders(w http.ResponseWriter, r *http.Request) {
	// content-security-policy
	w.Header().Set("Content-Security-Policy",
		"default-src 'self'; img-src 'self' data:; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline';")

	// access-control-allow-origin
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	// access-control-allow-headers
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// access-control-allow-methods
	w.Header().Set("Access-Control-Allow-Methods", "POST")
}

// GetUserID user id from username
func GetUserID(username string) int {
	// query db
	var id int
	db.DB.QueryRow("SELECT id FROM users WHERE username = $1;", username).Scan(&id)

	// return
	return id
}

// GetUsername username from user id
func GetUsername(id int) string {
	// query db
	var username string
	db.DB.QueryRow("SELECT username FROM users WHERE id = $1;", id).Scan(&username)

	// return
	return username
}
