package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

func (h *Handlers) LoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger.Debug(fmt.Sprintf("new incoming request > %s", r.RequestURI))
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Debug(fmt.Sprintf("< request took: %s", time.Since(start)))
	}

	return http.HandlerFunc(fn)
}
