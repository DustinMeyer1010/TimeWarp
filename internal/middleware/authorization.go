package middleware

import (
	"context"

	"net/http"
	"strings"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

type contextKey string

// Get Authorization token
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			if token == "" || !strings.HasPrefix(token, "bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claimsUnParsed, err := utils.VerifyAccessToken(token[:7])

			if err != nil {
				http.Error(w, "Failed to parse access token", http.StatusBadRequest)
				return
			}

			claims, err := types.CreateClaims(claimsUnParsed)

			if err != nil {
				http.Error(w, "Claims parse failed", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), contextKey("claims"), claims)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
}
