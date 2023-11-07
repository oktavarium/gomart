package auth

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
	"golang.org/x/crypto/pbkdf2"
)

var keyGenIter = 4096
var keyLength = 32

type Auth struct {
	key     string
	storage storager.Storager
	logger  logger.Logger
}

func NewAuth(logger logger.Logger, key string, storage storager.Storager) *Auth {
	return &Auth{key: key, storage: storage, logger: logger}
}

func (a *Auth) RegisterUser(ctx context.Context, user, password string) (string, error) {
	if !checkCredentials(user, password) {
		return "", ErrEmptyCredentials
	}

	exists, err := a.storage.UserExists(ctx, user)
	if err != nil {
		return "", fmt.Errorf("error on checking user existance: %w", err)
	}
	if exists {
		return "", ErrUserExists
	}

	hash, salt, err := generateHashAndSalt(user, password)
	if err != nil {
		return "", fmt.Errorf("error on generating user hash and salt: %w", err)
	}

	if err := a.storage.RegisterUser(ctx, user, hash, salt); err != nil {
		return "", fmt.Errorf("error on registering user in storage: %w", err)
	}

	token, err := a.generateToken(user)
	if err != nil {
		return "", fmt.Errorf("error on generating token: %w", err)
	}

	return token, nil
}

func (a *Auth) Authorize(ctx context.Context, user, password string) (string, error) {
	if !checkCredentials(user, password) {
		return "", ErrEmptyCredentials
	}

	exists, err := a.storage.UserExists(ctx, user)
	if err != nil {
		return "", fmt.Errorf("error on checking user existance: %w", err)
	}
	if exists {
		return "", ErrUserExists
	}

	storedHash, storedSalt, err := a.storage.UserHashAndSalt(ctx, user)
	if err != nil {
		return "", fmt.Errorf("error on getting user hash and salt from storage: %w", err)
	}

	saltBytes, err := base64ToBytes(storedSalt)
	if err != nil {
		return "", fmt.Errorf("error on decoding base64 salt: %w", err)
	}

	hashBytes := pbkdf2.Key([]byte(password), saltBytes, keyGenIter, keyLength, sha256.New)
	hash := bytesToBase64(hashBytes)

	if storedHash != hash {
		return "", ErrNotAuthorized
	}

	token, err := a.generateToken(user)
	if err != nil {
		return "", fmt.Errorf("error on generating token: %w", err)
	}

	return token, nil
}
