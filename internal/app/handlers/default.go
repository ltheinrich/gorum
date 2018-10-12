package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	// Handlers map
	Handlers = map[string]func(http.ResponseWriter, *http.Request){
		"login":    Login,
		"register": Register,
	}
)

// Header set
func Header(w http.ResponseWriter) {
	// set headers for CORS
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST") // GET, POST, PUT, DELETE

	// set header code
	w.WriteHeader(200)
}

// Read and unmarshal to map[string]interface{}
func Read(reader io.Reader, length int64) (map[string]interface{}, error) {
	var err error

	// read
	buffer := make([]byte, length)
	reader.Read(buffer)

	// unmarshal json
	var request map[string]interface{}
	err = json.Unmarshal(buffer, &request)
	if err != nil {
		return nil, err
	}

	// return map
	return request, nil
}

// Write and marshal from map[string]interface{}
func Write(writer io.Writer, response map[string]interface{}) error {
	var err error

	// marshal to byte slice
	var responseBytes []byte
	responseBytes, err = json.Marshal(&response)
	if err != nil {
		return err
	}

	// write
	_, err = writer.Write(responseBytes)
	if err != nil {
		return err
	}

	return nil
}

// Error print and write error
func Error(w http.ResponseWriter, err error) {
	log.Println(err)
	Write(w, map[string]interface{}{"error": err.Error()})
}

// ErrorWrite string
func ErrorWrite(w http.ResponseWriter, err string) {
	Write(w, map[string]interface{}{"error": err})
}
