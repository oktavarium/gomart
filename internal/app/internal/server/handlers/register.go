package handlers

import (
	"encoding/json"
	"net/http"
)

type user struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u user
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok, err := h.storage.UserExists(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if ok {
		w.WriteHeader(http.StatusConflict)
		return
	}

	token, err := h.storage.RegisterUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
