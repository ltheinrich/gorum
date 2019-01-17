package uuid

import (
	"crypto/rand"
	"fmt"
)

// Generate return UUID as string
func Generate() string {
	// make byte slice
	random := make([]byte, 16)

	// read from random
	rand.Read(random)

	// return uuid as string
	return fmt.Sprintf("%x-%x-%x-%x-%x", random[:4], random[4:6], random[6:8], random[8:10], random[10:16])
}
