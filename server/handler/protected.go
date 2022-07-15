package handler

import (
	"net/http"
)

func (h *Handler) protectedHandler(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Hello, I'm protected",
	})
}
