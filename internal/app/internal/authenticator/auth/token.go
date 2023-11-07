package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var tokenLifeTime = time.Hour * 24

type claims struct {
	jwt.RegisteredClaims
	User string
}

func (a *Auth) generateToken(user string) (string, error) {
	claims := claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenLifeTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(a.key))
	if err != nil {
		return "", fmt.Errorf("error on signing token: %w", err)
	}

	return ss, nil
}

func (a *Auth) GetUser(ctx context.Context, tokenString string) (string, error) {
	claims := &claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(a.key), nil
	})

	if err != nil {
		return "", fmt.Errorf("error on parsing token: %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("token is not valid")
	}

	exists, err := a.storage.UserExists(ctx, claims.User)
	if err != nil {
		return "", fmt.Errorf("error on checking user existance: %w", err)
	}

	if !exists {
		return "", ErrUserNotExists
	}

	return claims.User, nil
}
