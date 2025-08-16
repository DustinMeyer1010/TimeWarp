package server

import (
	"log"
	"net/http"
)

func Start() {
	router := createRouter()

	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
