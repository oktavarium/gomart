package app

import (
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/server"
)

func Run() error {
	c, err := loadConfig()
	if err != nil {
		err = fmt.Errorf("error on loading config: %w", err)
		logger.Error(err)
		return err
	}

	logger.Init(c.LogLevel)

	s, err := server.NewServer(c.DatabaseURI, c.AccrualAddress, c.SecretKey)
	if err != nil {
		err = fmt.Errorf("error on creating new server: %w", err)
		logger.Error(err)
		return err
	}

	return s.ListenAndServe(c.Address)
}
