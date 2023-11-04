package auth

type Storage interface {
	UserExists(string) (bool, error)
	RegisterUser(string, string, string) error
	UserHashAndSalt(string) (string, string, error)
}
