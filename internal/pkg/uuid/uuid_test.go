package uuid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	// generate uuid
	uuid := Generate()

	// validate uuid length
	if len(uuid) != 36 {
		// failed
		t.Error("Could not generate uuid, length is not 36")
	}
}
