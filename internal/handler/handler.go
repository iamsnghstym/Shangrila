package handler

// Handler holds the handler configuration
type Handler struct {
	*service.Service
}

// New creates a new router with pre-defined routing
func New(s *internal.Service) http.Handler {
	api := way.NewRouter()
	api.HandleFunc("POST", "/login", nil)
	api.HandleFunc("POST", "/register", nil)

	r := way.NewRouter()
	r.Handle("*", "/api", http.StripPrefix("/api", api))

	return r
}