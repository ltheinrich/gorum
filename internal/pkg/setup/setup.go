package setup

import (
	"log"
	"os"
	"path"
	"strings"
	"time"
)

// LogToFile set log output to file (append)
func LogToFile(fileName string) error {
	var err error

	// get log file
	logFile := time.Now().Format(fileName)

	// split directory from filename and create them
	if strings.Contains(fileName, "/") {
		directory, _ := path.Split(logFile)
		err = os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// open file and check for error
	var file *os.File
	file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// set file as output
	if err == nil {
		log.SetOutput(file)
	}

	return err
}
