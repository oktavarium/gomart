package accruals

import (
	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (a *Accruals) startUpdater(pointsCh <-chan model.Points) {
	go func() {
		for points := range pointsCh {

			select {
			case <-a.ctx.Done():
				return
			default:
				if points.Status == REGISTERED || points.Status == PROCESSING {
					points.Status = PROCESSING
				}
				err := a.storage.UpdateOrder(a.ctx, points.Order, points.Status, points.Accrual)
				if err != nil {
					a.logger.Error(err)
				}
			}
		}
	}()
}
