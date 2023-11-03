package handlers

import "net/http"

func (h *Handlers) SecurityMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

	}

	return http.HandlerFunc(fn)
}
