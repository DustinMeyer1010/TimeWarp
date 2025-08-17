package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var account types.Account

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		http.Error(w, "unable to parse body", http.StatusBadRequest)
		return
	}

	foundAccount, err := db.GetAccountByUsername(account.Username)

	fmt.Println(foundAccount)
	if err != nil {
		http.Error(w, "Failed to add to db", http.StatusBadRequest)
		return
	}

	if !foundAccount.CheckPassword(account) {
		http.Error(w, "invalid password", http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJWTAccessToken(account.Username)

	if err != nil {
		fmt.Println("Failed to generate token: ", err)
	}

	refreshToken, err := utils.GenerateRefreshToken(account.Username)

	if err != nil {
		fmt.Println("Failed to generate token: ", err)
	}

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
