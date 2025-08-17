package router

import (
	"github.com/DustinMeyer1010/TimeWarp/internal/handler"
	"github.com/gorilla/mux"
)

func habitRoutes(router *mux.Router) {
	router.HandleFunc("/create/habit", handler.CreateHabit).Methods("POST")
}
