package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

func (h *Handlers) NewOrder(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "text/plain" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("error on reading body: %w", err)
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
