package app

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	sp, err := newServiceProvider(ctx)
	if err != nil {
		err = fmt.Errorf("error on creating new service provider: %w", err)
		return err
	}

	return sp.Run()
}
