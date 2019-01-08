package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	// Handlers map
	Handlers = map[string]func(request map[string]interface{}, username string, auth bool) interface{}{
		"login":        Login,
		"register":     Register,
		"lang":         Lang,
		"conf":         Conf,
		"users":        Users,
		"user":         User,
		"editusername": EditUsername,
		"editpassword": EditPassword,
		"boards":       Boards,
		"board":        Board,
		"threads":      Threads,
		"thread":       Thread,
		"posts":        Posts,
		"newcaptcha":   NewCaptcha,
		"newthread":    NewThread,
		"deletethread": DeleteThread,
		"newpost":      NewPost,
		"deletepost":   DeletePost,
		"lastthreads":  LastThreads,
	}
)

// GenerateHandler add custom variables to handler
func GenerateHandler(handler func(request map[string]interface{}, username string, auth bool) interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// headers
		w.Header().Set("Content-Type", "application/json")
		SecurityHeaders(w, r)
		w.WriteHeader(200)

		// require POST request
		if r.Method != "POST" {
			w.Write([]byte{})
			return
		}

		// read request
		request := read(r.Body, r.ContentLength)

		// authenticate
		var auth bool
		username := GetString(request, "username")
		password := GetString(request, "password")
		if username != "" && password != "" {
			auth = login(username, password)
		}

		// handle
		response := handler(request, username, auth)

		// write response
		if err, isErr := response.(error); isErr {
			// write error string
			writeMap(w, map[string]interface{}{"error": err.Error()})
		} else if resp, isByte := response.([]byte); isByte {
			// write byte slice
			write(w, resp)
		} else if resp, isMap := response.(map[string]interface{}); isMap {
			// write response map
			writeMap(w, resp)
		} else {
			// unknown response type
			write(w, []byte(`{"error": "unknown response type"}`))
		}
	}
}

// read and unmarshal to map[string]interface{}
func read(reader io.Reader, length int64) map[string]interface{} {
	var err error

	// read
	buffer := make([]byte, length)
	reader.Read(buffer)

	// unmarshal json
	var request map[string]interface{}
	err = json.Unmarshal(buffer, &request)
	if err != nil {
		return map[string]interface{}{}
	}

	// return map
	return request
}

// write map marshalled from map[string]interface{}
func writeMap(writer io.Writer, response map[string]interface{}) error {
	var err error

	// marshal to byte slice
	var responseBytes []byte
	responseBytes, err = json.Marshal(&response)
	if err != nil {
		return err
	}

	// write
	return write(writer, responseBytes)
}

// write byte slice
func write(writer io.Writer, response []byte) error {
	var err error

	// write
	_, err = writer.Write(response)
	if err != nil {
		return err
	}

	return nil
}
