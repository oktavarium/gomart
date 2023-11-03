package auth

import "github.com/oktavarium/gomart/internal/app/internal/server/shared"

func GenToken(u shared.User) (string, error) {
	return "", nil
}

func Authorize(u shared.User) error {
	return nil
}

func GenUserInfo(u shared.User) (shared.UserInfo, error) {
	return shared.UserInfo{}, nil
}

func UserInfo(u shared.User) (shared.UserInfo, error) {
	return shared.UserInfo{}, nil
}
