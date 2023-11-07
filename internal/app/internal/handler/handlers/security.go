package handlers

import (
	"context"
	"net/http"
	"strings"
)

var tokenPrefix = "Bearer "

func (h *Handlers) SecurityMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				h.logger.Error(err)
			}
		}()

		tokenString := r.Header.Get("Authorization")
		token := strings.TrimPrefix(tokenString, tokenPrefix)
		login, err := h.authenticator.GetUser(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserLogin, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
