package handler

// Handler holds the handler configuration
type Handler struct {
	*service.Service
}

// New creates a new router with pre-defined routing
func New(s *internal.Service) http.Handler {
	h := Handler {s}

	api := way.NewRouter()
	api.HandleFunc("POST", "/login", h.login)
	api.HandleFunc("POST", "/register", h.createUser)

	r := way.NewRouter()
	r.Handle("*", "/api", http.StripPrefix("/api", api))

	return r
}