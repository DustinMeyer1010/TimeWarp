package handler

import (
	"net/http"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {

	token, ok := r.Context().Value("token").(string)

	if !ok {
		http.Error(w, "token context parse error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}
