package handlers

import (
	"github.com/nathannr/gorum/internal/pkg/config"
)

var (
	// Languages map
	Languages = map[string][]byte{}
)

// Lang handler
func Lang(data HandlerData) interface{} {
	// get language string
	language := data.Request.GetString("language")
	if language == "" {
		// custom language as hard fallback
		language = config.Get("public", "language")
	}

	// write language bytes
	return Languages[language]
}
