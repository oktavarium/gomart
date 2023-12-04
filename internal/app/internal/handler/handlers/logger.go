package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func (h *Handlers) LoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.logger.Debug(fmt.Sprintf(
			"new request > %s; Content-Type: %s;", r.RequestURI, r.Header.Get("Content-Type"),
		))

		start := time.Now()
		next.ServeHTTP(w, r)
		h.logger.Debug(fmt.Sprintf("< request took: %s", time.Since(start)))
	}

	return http.HandlerFunc(fn)
}
