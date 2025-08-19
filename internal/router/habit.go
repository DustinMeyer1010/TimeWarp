package router

import (
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/handler"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/gorilla/mux"
)

func habitRoutes(router *mux.Router) {
	router.Handle("/create/habit",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.CreateHabit),
			middleware.Authorization,
		),
	).Methods("POST")

	router.Handle("/all/habit/{id}",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.GetAllHabits),
			middleware.Authorization,
			middleware.VerifyIDWithToken,
		)).Methods("GET")

	router.Handle("/delete/habit/{id}",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.DeleteHabit),
			middleware.Authorization,
		)).Methods("DELETE")

}
