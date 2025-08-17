package handler

import (
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "No username in jwt.token", http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJWTAccessToken(username)

	if err != nil {
		http.Error(w, "Generation of access token failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}
