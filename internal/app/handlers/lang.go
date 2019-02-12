package handlers

import (
	"github.com/ltheinrich/gorum/internal/pkg/config"
)

var (
	// Languages map
	Languages = map[string][]byte{}
)

// Lang handler
func Lang(request map[string]interface{}, username string, auth bool) interface{} {
	// get language string
	language := GetString(request, "language")
	if language == "" {
		// custom language as hard fallback
		language = config.Get("public", "language")
	}

	// write language bytes
	return Languages[language]
}
