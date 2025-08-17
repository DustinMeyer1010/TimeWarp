package server

import (
	"log"
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/router"
)

func Start() {
	router := router.CreateRouter()

	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
