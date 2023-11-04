package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/auth"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.auth.RegisterUser(u.Login, u.Password)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrEmptyCredentials):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Is(err, auth.ErrUserExists):
			w.WriteHeader(http.StatusConflict)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
}
