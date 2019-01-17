package handlers

import (
	"errors"

	"github.com/ltheinrich/captcha"
	"github.com/ltheinrich/gorum/internal/pkg/config"
)

// NewCaptcha handler
func NewCaptcha(request map[string]interface{}, username string, auth bool) interface{} {
	// check if captcha enabled
	if config.Get("https", "captcha") != TRUE {
		// not enabled
		return errors.New("501")
	}

	// generate captcha and respond
	cap := captcha.New()
	return map[string]interface{}{"captcha": cap}
}
