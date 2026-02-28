package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Token struct {
	Username string `json:"username"`
	Due int64 `json:"expiration"`
	signature []byte `json:"-"`
}

func NewToken(username string) *Token {
	return &Token{Username: username}
}

func (t *Token) serialize() []byte {
	data, _ := json.Marshal(t)
	return data
}

func (t *Token) Issue(key []byte) string {
	t.Due = time.Now().Add(30*time.Minute).Unix()
	data := t.serialize()

	h := hmac.New(sha256.New, key)
	h.Write(data)
	signature := h.Sum(nil)

	return base64.StdEncoding.EncodeToString(data) + "." + base64.StdEncoding.EncodeToString(signature)
}

func (t *Token) Validate(key []byte) bool {
	data := t.serialize()

	h := hmac.New(sha256.New, key)
	h.Write(data)
	signature := h.Sum(nil)

	return time.Now().Before(time.Unix(t.Due, 0)) && hmac.Equal(t.signature, signature)
}

func ParseToken(token string) (*Token, error) {
	rawData, signature, found := strings.Cut(token, ".")
	if !found { return nil, fmt.Errorf("Invalid token") }

	var t Token
	var err error

	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil { return nil, fmt.Errorf("Invalid token") }

	if err = json.Unmarshal([]byte(data), &t); err != nil {
		return nil, fmt.Errorf("Invalid token")
	}

	t.signature, err = base64.StdEncoding.DecodeString(signature)
	if err != nil { return nil, fmt.Errorf("Invalid token") }

	return &t, nil
}
