package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

func (h *Handlers) Balance(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()

	user := r.Context().Value(UserLogin).(string)
	balance, err := h.orders.GetBalance(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
