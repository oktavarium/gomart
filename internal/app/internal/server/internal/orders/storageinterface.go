package orders

type Storage interface {
	NewOrder(string, string) error
}
