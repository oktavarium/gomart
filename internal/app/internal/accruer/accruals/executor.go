package accruals

import (
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (a *Accruals) startExecutor(orders <-chan string, bufferSize uint) <-chan model.Points {
	if bufferSize == 0 {
		bufferSize = defaultBufferize
	}

	outCh := make(chan model.Points, bufferSize)

	go func() {
		ticker := time.NewTicker(defaultRequestInterval)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case <-a.ctx.Done():
				close(outCh)
				return
			default:
				if order, ok := <-orders; ok {
					points, err := a.getPoints(order)
					if err != nil {
						a.logger.Error(err)
						continue
					}
					outCh <- points
				}
			}
		}
	}()

	return outCh
}
