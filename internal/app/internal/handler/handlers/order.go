package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/orderer/orders"
)

func (h *Handlers) MakeOrder(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			h.logger.Error(err)
		}
	}()

	if r.Header.Get("Content-Type") != "text/plain" {
		err = fmt.Errorf("wrong content type")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("error on reading body: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := r.Context().Value(UserLogin).(string)
	err = h.orderer.MakeOrder(r.Context(), user, string(order))
	if err != nil {
		switch {
		case errors.Is(err, orders.ErrWrongOrderNumber):
			w.WriteHeader(http.StatusUnprocessableEntity)
		case errors.Is(err, orders.ErrAnotherUserOrder):
			w.WriteHeader(http.StatusConflict)
		case errors.Is(err, orders.ErrLoadedOrder):
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
