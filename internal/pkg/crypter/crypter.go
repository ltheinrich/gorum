package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/sha3"
)

// Hash return SHA-3 512 hash string
func Hash(input string) string {
	// initialize hasher
	hasher := sha3.New512()

	// write string as byte slice to hasher
	hasher.Write([]byte(input))

	// return hash
	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateKey generate AES key
func GenerateKey(key string) []byte {
	// define length
	length := len(key)

	// switch for valid key lengths
	switch {
	// 32 byte key
	case length >= 32:
		return []byte(key[:32])

	// 24 byte key
	case length >= 24:
		return []byte(key[:24])

	// 16 byte key
	case length >= 16:
		return []byte(key[:16])
	}

	// return key or rerun with valid size
	return GenerateKey(key + "_gorum")
}

// Encrypt text with key
func Encrypt(text string, key []byte) (string, error) {
	var err error

	// text as byte slice
	plain := []byte(text)

	// initialize and check for error
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// generate cipher text
	cipherText := make([]byte, aes.BlockSize+len(plain))
	iv := cipherText[:aes.BlockSize]

	// read cipher text and check for error
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	// encrypt
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plain)

	// return encoded as string
	return base64.URLEncoding.EncodeToString(cipherText), err
}

// Decrypt text with key
func Decrypt(text string, key []byte) (string, error) {
	var err error

	// decode to byte slice
	var cipherText []byte
	cipherText, err = base64.URLEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	// initialize and check for error
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// generate cipher text
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// decrypt
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	// return as string
	return string(cipherText), nil
}
