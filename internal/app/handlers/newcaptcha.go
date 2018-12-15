package handlers

import (
	"github.com/dchest/captcha"
)

// NewCaptcha handler
func NewCaptcha(request map[string]interface{}, username string, auth bool) interface{} {
	// generate captcha and respond
	cap := captcha.New()
	return map[string]interface{}{"captcha": cap}
}
