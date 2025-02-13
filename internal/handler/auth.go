package handler

import (
	"encoding/json"
	"errors"
	"github.com/iamsnghstym/Shangrila/internal/service"
	"net/http"
)

// loginInput
type loginInput struct {
	Email string
}

// login handlers handles the logic for logging in users
func (h * handler) login(w http.ResponseWriter, r *http.Request) {
	var in loginInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.Login(r.Context(), in.Email)
	if errors.Is(err, service.ErrUserNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if errors.Is(err, service.ErrEmailInvalid) {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		respondError(w, err)
	}

	respond(w, out, http.StatusOK)
}
