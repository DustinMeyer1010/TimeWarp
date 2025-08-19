package router

import (
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/handler"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/gorilla/mux"
)

func accountRoutes(router *mux.Router) {
	router.HandleFunc("/account", handler.CreateAccount).Methods("POST")

	router.Handle("/account/{id}",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.DeleteAccount),
			middleware.Authorization,
		),
	).Methods("DELETE")

	router.HandleFunc("/account/login", handler.Login).Methods("POST")
}
