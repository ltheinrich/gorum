package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	// Handlers map
	Handlers = map[string]func(request map[string]interface{}, username string) interface{}{
		"login":    Login,
		"register": Register,
		"lang":     Lang,
		"conf":     Conf,
		"users":    Users,
	}
)

// GenerateHandler add custom variables to handler
func GenerateHandler(handler func(request map[string]interface{}, username string) interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// set headers for CORS
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST") // optional add: GET, PUT, DELETE

		// set header code
		w.WriteHeader(200)

		// write response
		response := handler(read(r.Body, r.ContentLength), "")
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
			write(w, []byte(`{"error": "unknwon response type"}`))
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
