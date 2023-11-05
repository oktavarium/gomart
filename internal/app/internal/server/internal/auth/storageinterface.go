package auth

import "context"

type Storage interface {
	UserExists(context.Context, string) (bool, error)
	RegisterUser(context.Context, string, string, string) error
	UserHashAndSalt(context.Context, string) (string, string, error)
}
