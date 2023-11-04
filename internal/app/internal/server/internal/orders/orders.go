package orders

type Orders struct {
	storage Storage
}

func NewOrders(storage Storage) *Orders {
	return &Orders{storage}
}

func (o *Orders) NewOrder(user, order string) error {
	order = compressOrderNumber(order)
	if !checkOrderNumber(order) {
		return ErrWrongOrderNum
	}

	return nil
}
