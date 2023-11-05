package accruals

import "errors"

var ErrNotRegistered = errors.New("order not registered")
var ErrTooManyRequests = errors.New("too many requests")
var ErrAccrualSystemError = errors.New("accrual system internal error")
