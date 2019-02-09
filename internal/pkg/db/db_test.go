package db

import (
	"testing"
)

func TestConnect(t *testing.T) {
	// connect to database
	err := Connect("::1", "5432", "disable", "gorum", "gorum", "gorum")

	// check for error
	if err != nil {
		// failed
		t.Errorf("Could not connect to database, %v\n", err)
	} else if DB == nil {
		// check if database variable is nil, failed
		t.Error("Could not connect to database, database variable is nil")
	}
}
