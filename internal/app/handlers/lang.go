package handlers

var (
	// Language file
	Language []byte
)

// Lang handler
func Lang(request map[string]interface{}, username string, auth bool) interface{} {
	// write language bytes
	return Language
}
