package configer

type Configer interface {
	Address() string
	DatabaseURI() string
	AccrualAddress() string
	LogLevel() string
	SecretKey() string
}
