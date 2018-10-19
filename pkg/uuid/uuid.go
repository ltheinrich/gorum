package uuid

import (
	"crypto/rand"
	"fmt"

	"github.com/lheinrichde/gorum/pkg/crypter"
)

// Generate return UUID as string (with hashing)
func Generate() string {
	// make byte slice
	random := make([]byte, 64)

	// read from random
	rand.Read(random)

	// hash
	hashed := crypter.HashByte(random)

	// return uuid as string
	return fmt.Sprintf("%x-%x-%x-%x-%x", hashed[2:6], hashed[8:10], hashed[12:14], hashed[19:21], hashed[25:31])
}

// Unsafe return UUID as string (without hashing)
func Unsafe() string {
	// make byte slice
	random := make([]byte, 16)

	// read from random
	rand.Read(random)

	// return uuid as string
	return fmt.Sprintf("%x-%x-%x-%x-%x", random[:4], random[4:6], random[6:8], random[8:10], random[10:16])
}
