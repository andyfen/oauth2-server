package handler

import (
	"encoding/json"
	"github.com/andyfen/oauth-server/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"net/http"
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
	r.Use(middleware.Logger)

	r.Get("/credentials", h.credentialsHandler)
	r.Get("/oauth2/token", h.tokenHandler)
	r.Post("/oauth2/token", h.tokenHandler)
	r.Get("/api/protected", auth.ValidateToken(h.protectedHandler, h.srv))

}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
