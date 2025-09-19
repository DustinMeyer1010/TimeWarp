package middleware

import (
	"context"
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
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

			ctx := context.WithValue(r.Context(), ContextKey("claims"), claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func GenerateJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			claims, ok := r.Context().Value(ContextKey("claims")).(models.Claims)

			if !ok {
				http.Error(w, "Failed to parse claims", http.StatusInternalServerError)
			}

			token, err := utils.GenerateJWTAccessToken(claims.ID, claims.Username)

			if err != nil {
				http.Error(w, "Generation of access token failed", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), ContextKey("token"), token)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
