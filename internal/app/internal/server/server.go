package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/auth"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/handlers"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/orders"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/storage"
)

var apiPath = "/api/user"

type Server struct {
	*chi.Mux
}

func NewServer(dbURI, accrualAddr, key string) (*Server, error) {
	server := &Server{chi.NewRouter()}
	storage, err := storage.NewStorage(dbURI)
	if err != nil {
		return nil, fmt.Errorf("error on creating storage: %w", err)
	}

	auth := auth.NewAuth(key, storage)
	orders := orders.NewOrders(storage)

	handlers := handlers.NewHandlers(auth, orders, accrualAddr)

	server.Route(apiPath, func(r chi.Router) {
		r.Use(handlers.LoggerMiddleware)

		r.Post("/register", handlers.Register)
		r.Post("/login", handlers.Login)

		r.Group(func(r chi.Router) {
			r.Use(handlers.SecurityMiddleware)

			r.Route("/orders", func(r chi.Router) {
				r.Post("/", handlers.NewOrder)
				r.Get("/", handlers.Orders)
			})

			r.Route("/balance", func(r chi.Router) {
				r.Get("/", handlers.Balance)
				r.Post("/withdraw", handlers.Withdraw)
			})

			r.Get("/withdrawals", handlers.Withdrawals)
		})

	})

	return server, nil
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}
