package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

func (h *Handlers) Orders(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()

	user := r.Context().Value(UserLogin).(string)
	orders, err := h.orders.GetOrders(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
