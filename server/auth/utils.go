package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func CreateClientID() string {
	str, _ := randomHex(32)
	return base64.URLEncoding.EncodeToString([]byte(str))
}

func CreateClientSecret() string {
	str, _ := randomHex(32)
	return base64.URLEncoding.EncodeToString([]byte(str))
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
