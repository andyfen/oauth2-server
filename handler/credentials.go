package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/models"
)

const (
	CLIENT_ID     = "000000"
	CLIENT_SECRET = "999999"
)

func (h *Handler) credentialsHandler(w http.ResponseWriter, r *http.Request) {

	err := h.clientStore.Set(CLIENT_ID, &models.Client{
		ID:     CLIENT_ID,
		Secret: CLIENT_SECRET,
		Domain: "http://localhost:8080",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"client_id":     CLIENT_ID,
		"client_secret": CLIENT_SECRET,
	})
}
