package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	// Handlers map
	Handlers = map[string]func(data HandlerData) interface{}{
		"login":           Login,
		"register":        Register,
		"lang":            Lang,
		"conf":            Conf,
		"users":           Users,
		"user":            User,
		"editusername":    EditUsername,
		"editpassword":    EditPassword,
		"deleteaccount":   DeleteAccount,
		"boards":          Boards,
		"board":           Board,
		"threads":         Threads,
		"thread":          Thread,
		"posts":           Posts,
		"newcaptcha":      NewCaptcha,
		"newthread":       NewThread,
		"deletethread":    DeleteThread,
		"newpost":         NewPost,
		"deletepost":      DeletePost,
		"lastthreads":     LastThreads,
		"lastuserthreads": LastUserThreads,
		"editthread":      EditThread,
		"editpost":        EditPost,
		"post":            Post,
		"footer":          Footer,
		"userdata":        UserData,
		"setuserdata":     SetUserData,
		"exportdata":      ExportData,
		"page":            Page,
	}
)

// GenerateHandler add custom variables to handler
func GenerateHandler(handler func(data HandlerData) interface{}) func(http.ResponseWriter, *http.Request) {
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

		// read request and defer body close
		request := read(r.Body, r.ContentLength)
		defer r.Body.Close()

		// authenticate
		var auth bool
		username := request.GetString("username")
		token := request.GetString("token")
		if username != "" && token != "" {
			auth = validateToken(username, token)
		}

		// handle
		response := handler(HandlerData{Request: request, Username: username, Authenticated: auth})

		// write response
		if err, isErr := response.(error); isErr {
			// print and write error string
			if !strings.HasPrefix(err.Error(), "4") {
				log.Printf("Handler error occured: %v\n", err.Error())
			}
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
func read(reader io.Reader, length int64) (request Request) {
	var err error

	// read
	buffer := make([]byte, length)
	io.ReadAtLeast(reader, buffer, int(length))

	// unmarshal json
	var requestMap map[string]interface{}
	err = json.Unmarshal(buffer, &requestMap)
	if err != nil {
		return Request{}
	}

	// return Request
	return Request{RequestMap: requestMap}
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
