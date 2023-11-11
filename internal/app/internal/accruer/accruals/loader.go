package accruals

import (
	"context"
	"fmt"
	"time"
)

func (a *Accruals) startLoader(ctx context.Context, bufferSize uint) <-chan string {
	unproccessedOrders, err := a.storage.OrdersByStatus(ctx, []string{NEW, PROCESSING})
	if err != nil {
		a.logger.Error(fmt.Errorf("error on getting orders by status: %w", err))
	}

	outCh := make(chan string, bufferSize)
	for _, order := range unproccessedOrders {
		outCh <- order
	}

	go func() {
		ticker := time.NewTicker(defaultRequestInterval)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case <-ctx.Done():
				close(outCh)
				return
			default:
				a.logger.Info("LOADING PROCESSING MESSAGES")
				unproccessedOrders, err := a.storage.OrdersByStatus(ctx, []string{PROCESSING})
				if err != nil {
					a.logger.Error(fmt.Errorf("error on getting orders by status: %w", err))
					continue
				}
				for _, order := range unproccessedOrders {
					outCh <- order
				}
			}
		}
	}()

	return outCh
}
