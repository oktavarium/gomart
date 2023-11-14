package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) GetBalance(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			h.logger.Error(err)
		}
	}()

	user := r.Context().Value(UserLogin).(string)
	balance, err := h.orderer.GetBalance(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
