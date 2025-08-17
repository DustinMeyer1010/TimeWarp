package middleware

import (
	"context"
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

func VerifyRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			refreshToken, err := r.Cookie("refresh_token")

			if err != nil {
				http.Error(w, "Could not find refresh token", http.StatusBadRequest)
				return
			}

			if refreshToken.Valid() != nil {
				http.Error(w, "Refresh token no longer valid", http.StatusBadRequest)
				return
			}

			claims, err := utils.VerifyRefreshToken(refreshToken.Value)

			if err != nil {
				http.Error(w, "Invalid refresh token", http.StatusBadRequest)
				return
			}

			username, ok := claims["username"].(string)

			if !ok {
				http.Error(w, "token parse error", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), contextKey("username"), username)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func GenerateJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("claims").(types.Claims)

			if !ok {
				http.Error(w, "Failed to parse claims", http.StatusInternalServerError)
			}

			token, err := utils.GenerateJWTAccessToken(claims.ID, claims.Username)

			if err != nil {
				http.Error(w, "Generation of access token failed", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), contextKey("token"), token)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
