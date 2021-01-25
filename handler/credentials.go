package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func makeUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("uuid")
	}

	return uuid.String()
}

func (h *Handler) credentialsHandler(w http.ResponseWriter, r *http.Request) {
	clientID := makeUUID()
	clientSecret := makeUUID()

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
