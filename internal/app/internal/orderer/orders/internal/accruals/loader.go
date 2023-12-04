package accruals

import (
	"context"
	"fmt"
	"time"
)

func (a *Accruals) startLoader(ctx context.Context, bufferSize uint) <-chan string {
	unprocessedOrders, err := a.storage.GetOrdersByStatus(ctx, []string{NEW, PROCESSING})
	if err != nil {
		a.logger.Error(fmt.Errorf("error on getting orders by status: %w", err))
	}

	outCh := make(chan string, bufferSize)
	for _, order := range unprocessedOrders {
		outCh <- order
	}

	go func() {
		ticker := time.NewTicker(defaultRequestInterval)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case <-a.ctx.Done():
				close(outCh)
				return
			default:
				unprocessedOrders, err := a.storage.GetOrdersByStatus(a.ctx, []string{PROCESSING})
				if err != nil {
					a.logger.Error(fmt.Errorf("error on getting orders by status: %w", err))
					continue
				}
				for _, order := range unprocessedOrders {
					outCh <- order
				}
			}
		}
	}()

	return outCh
}
