package handler

import (
	"context"
	"net/http"

	"github.com/andyfen/oauth-server/server/auth"
	"github.com/go-oauth2/oauth2/v4/models"
)

func (h *Handler) credentialsHandler(w http.ResponseWriter, r *http.Request) {
	clientID := auth.CreateClientID()
	clientSecret := auth.CreateClientSecret()

	info := &models.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: h.conf.DomainURL,
		UserID: "1_1",
	}

	err := h.clientStore.Create(context.Background(), info)
	/*
		err := h.clientStore.Set(clientID, &models.Client{
			ID:     clientID,
			Secret: clientSecret,
			Domain: h.conf.DomainURL,
		})
	*/
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
}
