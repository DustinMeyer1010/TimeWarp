package server

import (
	"log"
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/router"
	"github.com/gorilla/handlers"
)

func Start() {
	router := router.CreateRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),            // allow all origins
		handlers.AllowedMethods([]string{"GET", "POST"}),  // allowed methods
		handlers.AllowedHeaders([]string{"Content-Type"}), // allowed headers
	)(router)

	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: corsHandler,
	}

	log.Fatal(server.ListenAndServe())
}
