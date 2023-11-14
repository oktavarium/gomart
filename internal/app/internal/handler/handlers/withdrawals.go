package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) GetWithdrawals(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			h.logger.Error(err)
		}
	}()

	user := r.Context().Value(UserLogin).(string)
	withdrawls, err := h.orderer.GetWithdrawals(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(withdrawls) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(withdrawls)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
