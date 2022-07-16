package middleware

import (
	"net/http"

	"github.com/andyfen/oauth-server/server/contexts/token"
	"github.com/go-oauth2/oauth2/v4/server"
)

func ValidateOAuthToken(srv *server.Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenInfo, err := srv.ValidationBearerToken(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			ctx := token.NewContext(r.Context(), tokenInfo, err)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
