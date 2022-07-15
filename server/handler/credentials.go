package handler

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/models"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (h *Handler) credentialsHandler(w http.ResponseWriter, r *http.Request) {
	clientID, _ := randomHex(32)
	clientSecret, _ := randomHex(32)

	err := h.clientStore.Set(clientID, &models.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: "http://localhost:8080",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusOK, map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
}
