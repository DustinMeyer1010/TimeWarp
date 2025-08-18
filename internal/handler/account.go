package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/DustinMeyer1010/TimeWarp/internal/service"
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

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "unable to parse body", http.StatusBadRequest)
		return
	}

	if err := service.CreateAccount(account); err != nil {
		http.Error(w, "Failed to add to db", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.ContextKey("claims")).(types.Claims)

	if !ok {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}

	if err := service.DeleteAccount(id, claims.Username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
