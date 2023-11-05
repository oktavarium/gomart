package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

func (h *Handlers) Withdrawals(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()

	user := r.Context().Value(UserLogin).(string)
	withdrawls, err := h.orders.GetWithdrawals(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(&withdrawls)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
