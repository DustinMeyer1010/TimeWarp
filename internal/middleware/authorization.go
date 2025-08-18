package middleware

import (
	"context"

	"net/http"
	"strings"

	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

type ContextKey string

// Get Authorization token
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			if token == "" || !strings.HasPrefix(token, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, err := utils.VerifyAccessToken(token[7:])

			if err != nil {
				http.Error(w, "Failed to parse access token", http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), ContextKey("claims"), claims)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
}
