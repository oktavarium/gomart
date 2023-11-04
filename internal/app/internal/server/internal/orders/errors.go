package orders

import "errors"

var ErrWrongOrderNum = errors.New("wrong order number")
var ErrAnotherUserOrder = errors.New("order loaded used by another user")
var ErrLoadedOrder = errors.New("order already loaded")
