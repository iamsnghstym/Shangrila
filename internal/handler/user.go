package handler

import (
	"encoding/json"
	"errors"
	"github.com/iamsnghstym/Shangrila/internal/service"
	"net/http"
)

type CreateUserInput struct {
	Email, Username string
}

func (h * handler) createUser(w http.ResponseWriter, r *http.Request) {
	var in CreateUserInput
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.CreateUser(r.Context(), in.Email, in.Username)
	if errors.Is(err, service.ErrEmailInvalid) || errors.Is(err, service.ErrUsernameInvalid) {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if errors.Is(err, service.ErrEmailAlreadyTaken) || errors.Is(err, service.ErrUsernameAlreadyTaken) {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	
	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

