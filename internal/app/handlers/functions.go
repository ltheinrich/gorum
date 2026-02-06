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

// HandlerData structure
type HandlerData struct {
	Request       Request
	Username      string
	Authenticated bool
}

// Request structure
type Request struct {
	RequestMap map[string]interface{}
}

// GetString get string or empty string
func (request Request) GetString(name string) string {
	// cast and return value
	value, _ := request.RequestMap[name].(string)
	return value
}

// GetStringArray get string array or nil
func (request Request) GetStringArray(name string) []string {
	// cast
	raw, okRaw := request.RequestMap[name].([]interface{})

	// check if array
	if okRaw {
		// loop through string
		values := []string{}
		for _, value := range raw {
			// cast to string
			str, okStr := value.(string)

			// append
			if okStr {
				values = append(values, str)
			}
		}

		// return slice
		return values
	}

	// return
	return nil
}

// GetInt get int or 0
func (request Request) GetInt(name string) int {
	// cast and return value
	value, _ := request.RequestMap[name].(float64)

	// return int (rare, most common is float64 in JSON)
	if value == 0 {
		// return int
		i, _ := request.RequestMap[name].(int)
		return i
	}

	return int(value)
}

// GetBool get string or false
func (request Request) GetBool(name string) bool {
	// cast and return value
	value, _ := request.RequestMap[name].(bool)
	return value
}

// SecurityHeaders write CORS and CSP headers
func SecurityHeaders(w http.ResponseWriter, r *http.Request) {
	// content-security-policy
	w.Header().Set("Content-Security-Policy",
		"default-src 'self'; img-src 'self' data:; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; font-src 'self' https://fonts.gstatic.com/;")

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
