package crypter

import (
	"testing"
)

func TestHash(t *testing.T) {
	// hash
	hash := Hash("Test")

	// validate hash
	if hash != "301bb421c971fbb7ed01dcc3a9976ce53df034022ba982b97d0f27d48c4f03883aabf7c6bc778aa7c383062f6823045a6d41b8a720afbb8a9607690f89fbe1a7" {
		t.Error("Wrong hash value")
	}
}

func TestHashByte(t *testing.T) {
	// hash
	hash := HashByte([]byte("Test"))

	// loop through bytes
	for index, value := range []byte{48, 27, 180, 33, 201, 113, 251, 183, 237, 1, 220, 195, 169, 151, 108, 229, 61, 240, 52, 2, 43, 169, 130, 185, 125, 15, 39, 212, 140, 79, 3, 136, 58, 171, 247, 198, 188, 119, 138, 167, 195, 131, 6, 47, 104, 35, 4, 90, 109, 65, 184, 167, 32, 175, 187, 138, 150, 7, 105, 15, 137, 251, 225, 167} {
		// validate hash byte
		if value != hash[index] {
			// failed
			t.Error("Wrong hash value")
			break
		}
	}

}

func TestGenerateKey(t *testing.T) {
	// generate key
	key := GenerateKey("Hallo")

	// validate key
	if string(key) != "Hallo_gorum_goru" {
		// failed
		t.Error("Wrong key generation")
	}
}

func TestEncrypt(t *testing.T) {
	// encrypt
	encrypted, err := Encrypt("Das ist ein Test", GenerateKey("Hallo"))

	// check for error
	if err != nil {
		// failed
		t.Errorf("Could not encrypt, %v\n", err)
		return
	}

	// decrypt to verify
	var decrypted string
	decrypted, err = Decrypt(encrypted, GenerateKey("Hallo"))

	// check for error
	if err != nil {
		// failed
		t.Errorf("Could not encrypt, decryption failed: %v\n", err)
		return
	}

	// validate decryption
	if decrypted != "Das ist ein Test" {
		// failed
		t.Error("Could not encrypted, decryption failed: Wrong decrypted string")
	}
}

func TestDecrypt(t *testing.T) {
	// encrypt
	decrypted, err := Decrypt("kok1hUI21vkTZoIYjmNckHu8e5WxMV30TPzuWkuksZI=", GenerateKey("Hallo"))

	// check for error
	if err != nil {
		// failed
		t.Errorf("Could not decrypt, %v\n", err)
		return
	}

	// validate decryption
	if decrypted != "Das ist ein Test" {
		// failed
		t.Error("Could not decrypt, wrong decrypted string")
	}
}
