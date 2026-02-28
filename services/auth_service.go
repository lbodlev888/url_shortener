package services

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/lbodlev888/url_shortener/models"
	"golang.org/x/crypto/argon2"
)

const (
	timeCost uint32 = 3
	memoryCost uint32 = 64*1024
	threads uint8 = 4
	hashLen uint32 = 32
)

var path []byte
var key []byte

func init() {
	var err error
	path = make([]byte, 3)
	keyBase64 := os.Getenv("HMAC_SECRET")
	key, err = base64.StdEncoding.DecodeString(keyBase64)
	if err != nil { panic(err) }
}

func deriveKey(password []byte) (hash, salt string) {
	saltB := make([]byte, 16)
	rand.Read(saltB)

	hashB := argon2.Key(password, saltB, timeCost, memoryCost, threads, hashLen)

	hash = base64.StdEncoding.EncodeToString(hashB)
	salt = base64.StdEncoding.EncodeToString(saltB)

	return hash, salt
}

func checkKey(plainPassword, hashedPassword, salt []byte) bool {
	hashB := argon2.Key(plainPassword, salt, timeCost, memoryCost, threads, hashLen)
	return bytes.Equal(hashB, hashedPassword)
}

func ValidateToken(token string) bool {
	parsed, err := models.ParseToken(token)
	if err != nil {
		fmt.Printf("could not parse the token: %v\n", err)
		return false
	}

	return parsed.Validate(key)
}
