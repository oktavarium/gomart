package handler

import "net/http"

type Handler interface {
	LoggerMiddleware(http.Handler) http.Handler
	Login(http.ResponseWriter, *http.Request)
	NewOrder(http.ResponseWriter, *http.Request)
	Orders(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
	SecurityMiddleware(http.Handler) http.Handler
	Balance(http.ResponseWriter, *http.Request)
	Withdraw(http.ResponseWriter, *http.Request)
	Withdrawals(http.ResponseWriter, *http.Request)
}
