package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/authenticatorer/authenticator"
	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
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

	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.authenticatorer.RegisterUser(r.Context(), u.Login, u.Password)
	if err != nil {
		switch {
		case errors.Is(err, authenticator.ErrEmptyCredentials):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Is(err, authenticator.ErrUserExists):
			w.WriteHeader(http.StatusConflict)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
}
