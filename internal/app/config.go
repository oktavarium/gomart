package app

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env"
)

type config struct {
	Address        string `env:"RUN_ADDRESS"`            // адрес и порт запуска сервиса
	DatabaseURI    string `env:"DATABASE_URI"`           // адрес подключения к базе данных
	AccrualAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"` // адрес системы расчёта начислений
}

func loadConfig() (config, error) {
	var c config

	flag.StringVar(&c.Address, "a", "localhost:80", "listen address in notation address:port")
	flag.StringVar(&c.DatabaseURI, "d", "", "database connection string")
	flag.StringVar(&c.AccrualAddress, "r", "", "accrual system address")

	flag.Parse()

	if err := env.Parse(&c); err != nil {
		return c, fmt.Errorf("error on parsing env vars: %w", err)
	}

	if len(flag.Args()) > 0 {
		return c, fmt.Errorf("unrecognised flags")
	}

	switch {
	case len(c.DatabaseURI) == 0:
		return c, fmt.Errorf("empty database URI")
	case len(c.AccrualAddress) == 0:
		return c, fmt.Errorf("empty accrual system address")
	}

	return c, nil
}
