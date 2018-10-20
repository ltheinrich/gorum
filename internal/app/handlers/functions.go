package handlers

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
