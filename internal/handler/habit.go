package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	var habit types.Habit
	claims, ok := r.Context().Value(middleware.ContextKey("claims")).(types.Claims)

	if !ok {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&habit); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	habit.Account_id = claims.ID

	db.CreateHabit(habit)

}
