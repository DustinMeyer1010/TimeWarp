package router

import (
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	createAllRoutes(router,
		accountRoutes,
		habitRoutes,
		refreshRoutes,
	)

	return router
}

func createAllRoutes(router *mux.Router, routes ...func(*mux.Router)) {
	for _, route := range routes {
		route(router)
	}
}
