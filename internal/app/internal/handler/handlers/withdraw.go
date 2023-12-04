package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/model"
	"github.com/oktavarium/gomart/internal/app/internal/orderer/orders"
)

func (h *Handlers) Withdraw(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			h.logger.Error(err)
		}
	}()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var withdrawal model.Withdrawal
	err = json.NewDecoder(r.Body).Decode(&withdrawal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := r.Context().Value(UserLogin).(string)

	err = h.orderer.Withdraw(r.Context(), user, withdrawal.Order, withdrawal.Sum)
	if err != nil {
		switch {
		case errors.Is(err, orders.ErrWrongOrderNumber):
			w.WriteHeader(http.StatusUnprocessableEntity)
		case errors.Is(err, orders.ErrNotEnoughBalance):
			w.WriteHeader(http.StatusPaymentRequired)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}
