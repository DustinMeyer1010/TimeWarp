package router

import (
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/handler"
	"github.com/DustinMeyer1010/TimeWarp/internal/middleware"
	"github.com/gorilla/mux"
)

func refreshRoutes(router *mux.Router) {
	router.Handle("/refresh",
		middleware.ChainMiddleware(
			http.HandlerFunc(handler.RefreshToken),
			middleware.VerifyRefreshToken,
			middleware.GenerateJWTToken,
		),
	).Methods("GET")

}
