package uuid

import (
	"crypto/rand"
	"fmt"
)

// Generate return UUID as string
func Generate() string {
	// generate uuid bytes
	b1 := make([]byte, 4)
	b2 := make([]byte, 2)
	b3 := make([]byte, 2)
	b4 := make([]byte, 2)
	b5 := make([]byte, 4)
	rand.Read(b1)
	rand.Read(b2)
	rand.Read(b3)
	rand.Read(b4)
	rand.Read(b5)

	// return uuid as string
	return fmt.Sprintf("%x-%x-%x-%x-%x", b1, b2, b3, b4, b5)
}
