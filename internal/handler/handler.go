package handler

import (
	"github.com/iamsnghstym/Shangrila/internal/service"
	"github.com/matryer/way"
	"net/http"
)

// Handler holds the handler configuration
type handler struct {
	*service.Service
}

// New creates a new router with pre-defined routing
func New(s *service.Service) http.Handler {
	h := handler {s}

	api := way.NewRouter()
	api.HandleFunc("POST", "/login", h.login)
	api.HandleFunc("POST", "/register", h.createUser)

	r := way.NewRouter()
	r.Handle("*", "/api", http.StripPrefix("/api", api))

	return r
}