package orders

type Storage interface {
	NewOrder(string, string) error
	GetUserByOrder(string) (string, error)
	CreateOrder(string, string) error
}
