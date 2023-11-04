package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

var tokenPrefix = "Bearer "

func (h *Handlers) SecurityMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				logger.Error(err)
			}
		}()

		tokenString := r.Header.Get("Authorization")
		token := strings.TrimPrefix(tokenString, tokenPrefix)
		login, err := h.auth.GetUser(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserLogin, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
