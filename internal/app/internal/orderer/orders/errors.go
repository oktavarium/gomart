package orders

import "errors"

var ErrWrongOrderNumber = errors.New("wrong order number")
var ErrAnotherUserOrder = errors.New("order loaded used by another user")
var ErrLoadedOrder = errors.New("order already loaded")
var ErrNotEnoughBalance = errors.New("not enough balance")
