package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/oktavarium/gomart/internal/app/internal/provider"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	sp, err := provider.NewServiceProvider(ctx)
	if err != nil {
		err = fmt.Errorf("error on creating new service provider: %w", err)
		return err
	}

	return sp.Run()
}
