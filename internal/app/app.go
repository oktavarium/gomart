package app

import (
	"context"
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/provider"
)

func Run() error {
	ctx := context.Background()
	sp, err := provider.NewServiceProvider(ctx)
	if err != nil {
		err = fmt.Errorf("error on creating new service provider: %w", err)
		return err
	}

	return sp.Run()
}
