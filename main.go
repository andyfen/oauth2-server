package main

import (
	"github.com/andyfen/oauth-server/auth"
	"github.com/andyfen/oauth-server/config"
	"github.com/andyfen/oauth-server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	config := config.New()

	clientStore := auth.NewClientStore()
	authManager := auth.NewAuthManager(config, clientStore)
	authServer := auth.NewAuthServer(authManager)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	handler := handler.New(authServer, authManager, clientStore)
	handler.Register(r)

	http.ListenAndServe(":8080", r)
}
