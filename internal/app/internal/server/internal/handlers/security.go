package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

var tokenPrefix = "Bearer "

func (h *Handlers) SecurityMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token := strings.TrimPrefix(tokenString, tokenPrefix)
		login, err := h.auth.GetUser(token)
		if err != nil {
			err = fmt.Errorf("error on validating token: %w", err)
			logger.Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserLogin, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
