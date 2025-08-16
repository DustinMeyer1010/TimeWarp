package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account types.Account

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		http.Error(w, "unable to parse body", http.StatusBadRequest)
		return
	}

	if err = db.CreateAccount(account); err != nil {
		http.Error(w, "Failed to add to db", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var account types.Account

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		http.Error(w, "unable to parse body", http.StatusBadRequest)
		return
	}

	foundAccount, err := db.GetAccountByUsername(account.Username)
	if err != nil {
		http.Error(w, "Failed to add to db", http.StatusBadRequest)
		return
	}

	fmt.Println(foundAccount.Username, foundAccount.Password)
	fmt.Println(account.Username, account.Password)

	if !foundAccount.Verify(&account) {
		http.Error(w, "invalid password", http.StatusBadRequest)
		return
	}
}
