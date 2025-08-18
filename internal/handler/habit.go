package handler

import (
	"fmt"
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.ContextKey("claims")).(types.Claims)

	if !ok {
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	fmt.Println(claims.ID)
}
