package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	foundAccount, err := db.GetAccountByUsername(account.Username)

	if err != nil {
		http.Error(w, "Accouont not found", http.StatusBadRequest)
		return
	}

	if !foundAccount.CheckPassword(account) {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	token, _ := utils.GenerateJWTAccessToken(foundAccount.ID, account.Username) // errors are ignored since this would be configuration error

	refreshToken, _ := utils.GenerateRefreshToken(foundAccount.ID, account.Username) // errors are ignored since this would be configuration error

	db.AddRefreshToken(foundAccount.ID, refreshToken)

	refreshTokenCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	}

	http.SetCookie(w, &refreshTokenCookie)
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)

}
