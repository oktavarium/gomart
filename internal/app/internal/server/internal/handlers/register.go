package handlers

import (
	"encoding/json"
	"fmt"
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

	exists, err := h.storage.UserExists(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusConflict)
		return
	}

	u.Info, err = auth.GenUserInfo(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Println(err)
	}

	if err := h.storage.RegisterUser(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	token, err := auth.GenToken(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
}
