package router

import (
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/handler"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/gorilla/mux"
)

func habitRoutes(router *mux.Router) {

	router.Handle("/habit",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.CreateHabit),
			middleware.Authorization,
		),
	).Methods("POST")

	router.Handle("/habits/{id}",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.GetAllHabits),
			middleware.Authorization,
			middleware.VerifyIDWithToken,
		)).Methods("GET")

	router.Handle("/habit/time/{id}",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.DeleteHabitWithTime),
			middleware.Authorization,
		)).Methods("DELETE")

	router.Handle("/habit/{id}",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.DeleteHabitWithouttime),
			middleware.Authorization,
		)).Methods("DELETE")

}
