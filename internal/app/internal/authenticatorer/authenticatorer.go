package authenticatorer

import "context"

type Authenticatorer interface {
	RegisterUser(context.Context, string, string) (string, error)
	Authorize(context.Context, string, string) (string, error)
	GetUser(context.Context, string) (string, error)
}
