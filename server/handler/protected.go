package handler

import (
	"net/http"
	"time"

	"github.com/andyfen/oauth-server/server/contexts/token"
)

func (h *Handler) protectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenInfo, err := token.FromContext(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, ":(")
	}

	respondwithJSON(w, http.StatusOK, map[string]interface{}{
		"message":            "Hello, I'm protected",
		"expires_in":         int64(tokenInfo.GetAccessCreateAt().Add(tokenInfo.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":          tokenInfo.GetClientID(),
		"user_id":            tokenInfo.GetUserID(),
		"redirect_uri":       tokenInfo.GetRedirectURI(),
		"scope":              tokenInfo.GetScope(),
		"access_token":       tokenInfo.GetAccess(),
		"access_create_at":   tokenInfo.GetAccessCreateAt(),
		"access_expires_in":  tokenInfo.GetAccessExpiresIn(),
		"refresh_token":      tokenInfo.GetRefresh(),
		"refresh_create_at":  tokenInfo.GetRefreshCreateAt(),
		"refresh_expires_in": tokenInfo.GetRefreshExpiresIn(),
	})
}
