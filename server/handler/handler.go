package handler

import (
	"github.com/andyfen/oauth-server/server/middleware"
	"github.com/go-chi/chi"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

// Handler route handler
type Handler struct {
	srv         *server.Server
	manager     *manage.Manager
	clientStore *store.ClientStore
}

// New handler
func New(srv *server.Server, manager *manage.Manager, clientStore *store.ClientStore) *Handler {
	return &Handler{
		srv:         srv,
		manager:     manager,
		clientStore: clientStore,
	}
}

// Register - registers all the routes
func (h *Handler) Register(r *chi.Mux) {

	r.Get("/credentials", h.credentialsHandler)

	r.Route("/oauth2", func(r chi.Router) {
		r.Get("/token", h.tokenHandler)
		r.Post("/token", h.tokenHandler)
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.ValidateOAuthToken(h.srv))

		r.Get("/protected", h.protectedHandler)
		r.Post("/protected", h.protectedHandler)
	})
}
