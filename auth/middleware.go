package auth

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-oauth2/oauth2/v4/server"
)

// ValidateToken ...
func ValidateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.WithFields(log.Fields{
			"scope":    token.GetScope(),
			"clientID": token.GetClientID(),
			"userID":   token.GetUserID(),
			"code":     token.GetCode(),
			"access":   token.GetAccess(),
			"refresh":  token.GetRefresh(),
		}).Info("got token")

		// access=Z24RFU_ZMWGFT62IJES3DQ clientID=242da7fb code= refresh= scope=all userID=

		f.ServeHTTP(w, r)
	})
}
