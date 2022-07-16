package server

import (
	"net/http"
	"time"

	"github.com/andyfen/oauth-server/server/auth"
	"github.com/andyfen/oauth-server/server/config"
	"github.com/andyfen/oauth-server/server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New(conf *config.Config) (*http.Server, error) {
	clientStore := auth.NewClientStore()
	authManager := auth.NewAuthManager(conf, clientStore)
	authServer := auth.NewAuthServer(authManager)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	handler := handler.New(authServer, authManager, clientStore)
	handler.Register(r)

	srv := &http.Server{
		Addr:         ":" + conf.PORT,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return srv, nil
}
