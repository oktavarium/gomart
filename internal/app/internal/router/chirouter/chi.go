package chirouter

import (
	"context"
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
		addr:   addr,
		logger: logger,
	}

	server.Route(apiPath, func(r chi.Router) {
		r.Use(handler.LoggerMiddleware)

		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)

		r.Group(func(r chi.Router) {
			r.Use(handler.SecurityMiddleware)

			r.Route("/orders", func(r chi.Router) {
				r.Post("/", handler.NewOrder)
				r.Get("/", handler.Orders)
			})

			r.Route("/balance", func(r chi.Router) {
				r.Get("/", handler.Balance)
				r.Post("/withdraw", handler.Withdraw)
			})

			r.Get("/withdrawals", handler.Withdrawals)
		})

	})

	return server
}

func (s *ChiRouter) Run() error {
	server := &http.Server{Addr: s.addr, Handler: s}
	go func() {
		<-s.ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	return server.ListenAndServe()
}
