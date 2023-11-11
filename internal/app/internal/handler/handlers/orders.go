package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) Orders(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			h.logger.Error(err)
		}
	}()

	user := r.Context().Value(UserLogin).(string)
	orders, err := h.orderer.Orders(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
