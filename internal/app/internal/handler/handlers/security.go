package handlers

import (
	"context"
	"net/http"
)

func (h *Handlers) SecurityMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				h.logger.Error(err)
			}
		}()

		token := r.Header.Get("Authorization")
		login, err := h.authenticatorer.GetUser(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserLogin, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
