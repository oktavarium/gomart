package authenticatorer

import "errors"

var ErrUserExists = errors.New("user exists")
var ErrUserNotExists = errors.New("user not exists")
var ErrEmptyCredentials = errors.New("empty credentials")
var ErrNotAuthorized = errors.New("not authorized")
