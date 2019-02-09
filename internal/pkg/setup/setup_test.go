package setup

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestLogToFile(t *testing.T) {
	// log to file
	err := LogToFile("/tmp/gorum-test-log_file.log." + strconv.Itoa(rand.Intn(1000)))

	// check for error
	if err != nil {
		// failed
		t.Errorf("Could not log to file, %v\n", err)
	}
}
