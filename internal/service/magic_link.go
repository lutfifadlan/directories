package service

import (
	"crypto/rand"
	"encoding/base64"
)

const TokenLength = 32

func GenerateToken() (string, error) {
	bytes := make([]byte, TokenLength)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
