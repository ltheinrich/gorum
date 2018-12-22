package handlers

import (
	"net/http"

	"github.com/ltheinrich/gorum/pkg/db"
)

// GetString get string or empty string
func GetString(request map[string]interface{}, name string) string {
	// cast
	value, ok := request[name].(string)

	// return value
	if ok {
		return value
	}

	// return empty string
	return ""
}

// GetStringArray get string array or nil
func GetStringArray(request map[string]interface{}, name string) []string {
	// cast
	raw, okRaw := request[name].([]interface{})

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
func GetInt(request map[string]interface{}, name string) int {
	// cast
	value, ok := request[name].(int)

	// check exists
	if ok {
		// return int
		return value
	}

	// cast float
	var float float64
	float, ok = request[name].(float64)

	// check float exists
	if ok {
		// return float as int
		return int(float)
	}

	// return 0
	return 0
}

// GetBool get string or false
func GetBool(request map[string]interface{}, name string) bool {
	// cast
	value, ok := request[name].(bool)

	// return value
	if ok {
		return value
	}

	// return false
	return false
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
