package handlers

import "net/http"

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
		"default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline';")

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
