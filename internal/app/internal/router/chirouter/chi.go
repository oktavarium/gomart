package chirouter

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oktavarium/gomart/internal/app/internal/handler"
	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

var apiPath = "/api/user"

type ChiRouter struct {
	*chi.Mux
	ctx    context.Context
	addr   string
	logger logger.Logger
}

func NewRouter(
	ctx context.Context,
	logger logger.Logger,
	addr string,
	handler handler.Handler,
) *ChiRouter {

	server := &ChiRouter{
		Mux:    chi.NewRouter(),
		ctx:    ctx,
		addr:   addr,
		logger: logger,
	}

	server.Get("/ping", handler.Ping)

	server.Route(apiPath, func(r chi.Router) {
		r.Use(handler.LoggerMiddleware)

		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)

		r.Group(func(r chi.Router) {
			r.Use(handler.SecurityMiddleware)

			r.Route("/orders", func(r chi.Router) {
				r.Post("/", handler.MakeOrder)
				r.Get("/", handler.GetOrders)
			})

			r.Route("/balance", func(r chi.Router) {
				r.Get("/", handler.GetBalance)
				r.Post("/withdraw", handler.Withdraw)
			})

			r.Get("/withdrawals", handler.GetWithdrawals)
		})

	})

	return server
}

func (s *ChiRouter) Run() error {
	server := &http.Server{Addr: s.addr, Handler: s}
	go func() {
		<-s.ctx.Done()
		s.logger.Debug("context is done... exiting")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			s.logger.Error(fmt.Errorf("error on server shutdown: %w", err))
		}
	}()

	return server.ListenAndServe()
}
