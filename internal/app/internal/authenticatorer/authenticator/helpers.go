package authenticator

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func generateHashAndSalt(user, password string) (string, string, error) {
	saltBytes := make([]byte, 16)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", "", fmt.Errorf("error on reading random bytes: %w", err)
	}

	hashBytes := pbkdf2.Key([]byte(password), saltBytes, keyGenIter, keyLength, sha256.New)

	hash := bytesToBase64(hashBytes)
	salt := bytesToBase64(saltBytes)

	return hash, salt, nil
}

func bytesToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func base64ToBytes(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func checkCredentials(creds ...string) bool {
	for _, c := range creds {
		if len(c) == 0 {
			return false
		}
	}

	return true
}
