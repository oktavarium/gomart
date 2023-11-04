package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

func (h *Handlers) Withdraw(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var withdraw model.Withdraw
	err = json.NewDecoder(r.Body).Decode(&withdraw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := r.Context().Value(UserLogin).(string)

	h.orders.Withdraw(user, withdraw)
}
