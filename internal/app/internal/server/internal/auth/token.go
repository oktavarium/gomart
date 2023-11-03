package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

var tokenLifeTime = time.Hour * 24
var SECRET_KEY = []byte("test")

type claims struct {
	jwt.RegisteredClaims
	UserLogin string
}

func GenToken(u model.User) (string, error) {
	if len(u.Login) == 0 {
		return "", fmt.Errorf("empty user login")
	}
	claims := claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenLifeTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		u.Login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", fmt.Errorf("error on signing token: %w", err)
	}

	return ss, nil
}

func GetLogin(tokenString string) (string, error) {
	claims := claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return "", fmt.Errorf("error on parsing token: %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("token is not valid")
	}

	return claims.UserLogin, nil
}
