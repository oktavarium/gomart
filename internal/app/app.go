package app

import (
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/provider"
)

func Run() error {
	sp, err := provider.NewServiceProvider()
	if err != nil {
		err = fmt.Errorf("error on creating new service provider: %w", err)
		return err
	}

	return sp.Run()
}
