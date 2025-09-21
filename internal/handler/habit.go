package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/DustinMeyer1010/TimeWarp/internal/service"
	"github.com/gorilla/mux"
)

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	var habit models.Habit

	claims := r.Context().Value(middleware.ContextKey("claims")).(models.Claims)

	if err := json.NewDecoder(r.Body).Decode(&habit); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if _, err := service.CreateHabit(habit, claims); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetAllHabits(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.ContextKey("claims")).(models.Claims)

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

func DeleteHabitWithTime(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.ContextKey("claims")).(models.Claims)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid habit id", http.StatusBadRequest)
		return
	}

	if _, err := service.DeleteHabitWithTime(id, claims.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func DeleteHabitWithouttime(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.ContextKey("claims")).(models.Claims)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid habit id", http.StatusBadRequest)
		return
	}

	if _, err := service.DeleteHabitWithoutTime(id, claims.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
