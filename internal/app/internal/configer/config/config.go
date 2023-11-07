package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env"
)

type config struct {
	address        string `env:"RUN_ADDRESS"`                      // адрес и порт запуска сервиса
	databaseURI    string `env:"DATABASE_URI"`                     // адрес подключения к базе данных
	accrualAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`           // адрес системы расчёта начислений
	logLevel       string `env:"LOG_LEVEL"`                        // уровень логирования сервиса
	secretKey      string `env:"SECRET_KEY" envDefault:"test_key"` // ключ для подписывания токена
}

func (c *config) Address() string {
	return c.address
}

func (c *config) DatabaseURI() string {
	return c.databaseURI
}

func (c *config) AccrualAddress() string {
	return c.accrualAddress
}

func (c *config) LogLevel() string {
	return c.logLevel
}

func (c *config) SecretKey() string {
	return c.secretKey
}

func NewConfig() (*config, error) {
	var c config

	flag.StringVar(&c.address, "a", "localhost:80", "listen address in notation address:port")
	flag.StringVar(&c.databaseURI, "d", "", "database connection string")
	flag.StringVar(&c.accrualAddress, "r", "", "accrual system address")
	flag.StringVar(&c.logLevel, "l", "debug", "logger debug level")

	flag.Parse()

	if err := env.Parse(&c); err != nil {
		return nil, fmt.Errorf("error on parsing env vars: %w", err)
	}

	if len(flag.Args()) > 0 {
		return nil, fmt.Errorf("unrecognised flags")
	}

	if len(c.accrualAddress) == 0 {
		return nil, fmt.Errorf("empty accrual system address")
	}

	return &c, nil
}
