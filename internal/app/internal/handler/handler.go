package handler

import "net/http"

type Handler interface {
	LoggerMiddleware(http.Handler) http.Handler
	Login(http.ResponseWriter, *http.Request)
	MakeOrder(http.ResponseWriter, *http.Request)
	GetOrders(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
	SecurityMiddleware(http.Handler) http.Handler
	GetBalance(http.ResponseWriter, *http.Request)
	Withdraw(http.ResponseWriter, *http.Request)
	GetWithdrawals(http.ResponseWriter, *http.Request)
}
