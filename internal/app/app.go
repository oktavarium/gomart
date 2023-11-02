package app

import (
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/server"
)

func Run() error {
	c, err := loadConfig()
	if err != nil {
		return fmt.Errorf("error on loading config: %w", err)
	}

	s, err := server.NewServer(c.DatabaseURI, c.AccrualAddress)
	if err != nil {
		return fmt.Errorf("error on creating new server: %w", err)
	}

	return s.ListenAndServe(c.Address)
}
