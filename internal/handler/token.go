package handler

import (
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {

	token, ok := r.Context().Value(middleware.ContextKey("token")).(string)

	if !ok {
		http.Error(w, "token context parse error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}
