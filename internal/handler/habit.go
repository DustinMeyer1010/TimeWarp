package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/DustinMeyer1010/TimeWarp/internal/service"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/gorilla/mux"
)

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	var habit types.Habit

	claims := r.Context().Value(middleware.ContextKey("claims")).(types.Claims)

	if err := json.NewDecoder(r.Body).Decode(&habit); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if err := service.CreateHabit(habit, claims); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetAllHabits(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.ContextKey("claims")).(types.Claims)

	if !ok {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	habits, err := db.GetAllHabitsForUser(claims.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnBody, err := json.Marshal(habits)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(returnBody)
	w.WriteHeader(http.StatusOK)
}

func DeleteHabit(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.ContextKey("claims")).(types.Claims)

	if !ok {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid habit id", http.StatusBadRequest)
		return
	}

	if err := service.DeleteHabit(id, claims.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func TEST(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	db.CheckForCompletions(1, now)

}
