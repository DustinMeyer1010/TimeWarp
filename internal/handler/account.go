package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	db.Init()
	code := m.Run()

	os.Exit(code)
}

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

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	token_username := r.Context().Value("username").(string)

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}

	foundAccount, err := db.GetAccountByID(id)

	if err != nil {
		http.Error(w, "Account not found", http.StatusBadRequest)
		return
	}

	if token_username != foundAccount.Username {
		http.Error(w, "Unauthorized usernaem don't match", http.StatusUnauthorized)
		return
	}

	err = db.DeleteAccount(foundAccount.Username)

	if err != nil {
		http.Error(w, "Account not deleted", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
