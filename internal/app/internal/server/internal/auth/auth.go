package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
	"golang.org/x/crypto/pbkdf2"
)

var keyGenIter = 4096
var keyLength = 32

func Authorize(u model.User) (bool, error) {
	if len(u.Login) == 0 || len(u.Password) == 0 {
		return false, fmt.Errorf("empty credentials")
	}

	if len(u.Info.Hash) == 0 || len(u.Info.Salt) == 0 {
		return false, fmt.Errorf("empty user info")
	}

	saltBytes, err := base64ToBytes(u.Info.Salt)
	if err != nil {
		return false, fmt.Errorf("error on decoding base64 salt: %w", err)
	}

	hashBytes := pbkdf2.Key([]byte(u.Password), saltBytes, keyGenIter, keyLength, sha256.New)
	hash := bytesToBase64(hashBytes)

	return hash == u.Info.Hash, nil
}

func GenUserInfo(u model.User) (info model.UserInfo, err error) {
	if len(u.Login) == 0 || len(u.Password) == 0 {
		return u.Info, fmt.Errorf("empty credentials")
	}

	saltBytes := make([]byte, 16)
	if _, err := rand.Read(saltBytes); err != nil {
		return u.Info, fmt.Errorf("error on reading random bytes: %w", err)
	}

	hashBytes := pbkdf2.Key([]byte(u.Password), saltBytes, keyGenIter, keyLength, sha256.New)

	info.Hash = bytesToBase64(hashBytes)
	info.Salt = bytesToBase64(saltBytes)

	return info, nil
}

func bytesToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func base64ToBytes(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
