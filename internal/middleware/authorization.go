package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

type contextKey string

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			if token == "" || !strings.HasPrefix(token, "bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, err := utils.VerifyAccessToken(token[:7])

			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			username, ok := claims["username"].(string)

			if !ok {
				http.Error(w, "bad token claims", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), contextKey("username"), username)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
}
