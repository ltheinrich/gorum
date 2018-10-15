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
		"lang":     Lang,
		"conf":     Conf,
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
func Read(reader io.Reader, length int64) map[string]interface{} {
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

// Code write error string
func Code(w http.ResponseWriter, err string) {
	Write(w, map[string]interface{}{"error": err})
}

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
	value, ok := request[name].([]string)

	// return value
	if ok {
		return value
	}

	// return empty string
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
