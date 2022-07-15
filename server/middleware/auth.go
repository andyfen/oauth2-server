package middleware

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
)

func ValidateOAuthToken(srv *server.Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := srv.ValidationBearerToken(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
