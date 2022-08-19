package handler

import (
	"context"
	"net/http"

	"github.com/andyfen/oauth-server/server/auth"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
)

func (h *Handler) credentialsHandler(w http.ResponseWriter, r *http.Request) {
	clientID := auth.CreateClientID()
	clientSecret := auth.CreateClientSecret()

	userId, _ := uuid.NewUUID()

	info := &models.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: h.conf.DomainURL,
		UserID: userId.String(),
	}

	err := h.clientStore.Create(context.Background(), info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
}
