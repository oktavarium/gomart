package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/caarlos0/env"
)

type config struct {
	CAddress        string `env:"RUN_ADDRESS"`                      // адрес и порт запуска сервиса
	CDatabaseURI    string `env:"DATABASE_URI"`                     // адрес подключения к базе данных
	CAccrualAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`           // адрес системы расчёта начислений
	CLogLevel       string `env:"LOG_LEVEL"`                        // уровень логирования сервиса
	CSecretKey      string `env:"SECRET_KEY" envDefault:"test_key"` // ключ для подписывания токена
	CBufferSize     uint   `env:"BUFFER_SIZE" envDefault:"100"`     // размер буфферов для обработки
}

func (c *config) Address() string {
	return c.CAddress
}

func (c *config) DatabaseURI() string {
	return c.CDatabaseURI
}

func (c *config) AccrualAddress() string {
	return c.CAccrualAddress
}

func (c *config) LogLevel() string {
	return c.CLogLevel
}

func (c *config) SecretKey() string {
	return c.CSecretKey
}

func (c *config) BufferSize() uint {
	return c.CBufferSize
}

func NewConfig() (*config, error) {
	var c config

	flag.StringVar(&c.CAddress, "a", "localhost:80", "listen address in notation address:port")
	flag.StringVar(&c.CDatabaseURI, "d", "", "database connection string")
	flag.StringVar(&c.CAccrualAddress, "r", "", "accrual system address")
	flag.StringVar(&c.CLogLevel, "l", "debug", "logger debug level")
	flag.UintVar(&c.CBufferSize, "b", 100, "default buffer size")

	flag.Parse()

	if err := env.Parse(&c); err != nil {
		return nil, fmt.Errorf("error on parsing env vars: %w", err)
	}

	if len(flag.Args()) > 0 {
		return nil, fmt.Errorf("unrecognised flags")
	}

	if len(c.CAccrualAddress) == 0 {
		return nil, fmt.Errorf("empty accrual system address")
	}

	if !strings.Contains(c.AccrualAddress(), "http") {
		c.CAccrualAddress = "http://" + c.CAccrualAddress
	}

	return &c, nil
}
