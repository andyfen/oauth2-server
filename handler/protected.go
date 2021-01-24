package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) protectedHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"message": "Hello, I'm protected",
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(data)
}
