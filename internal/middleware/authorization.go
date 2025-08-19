package middleware

import (
	"context"
	"strconv"

	"net/http"
	"strings"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
	"github.com/gorilla/mux"
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

func VerifyIDWithToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(ContextKey("claims")).(types.Claims)

			if !ok {
				http.Error(w, "invalid token", http.StatusBadRequest)
				return
			}

			vars := mux.Vars(r)
			id, err := strconv.Atoi(vars["id"])

			if err != nil {
				http.Error(w, "invalid id", http.StatusBadRequest)
				return
			}

			if id != claims.ID {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)

		})
}
