package handler

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

	fmt.Println(clientID)
	fmt.Println(clientSecret)

	err := h.clientStore.Set(clientID, &models.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: "http://localhost:8080",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
}
