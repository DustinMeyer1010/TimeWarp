package server

import (
	"net/http"

	"github.com/DustinMeyer1010/TimeWarp/internal/handler"
)

func createRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create/account", handler.CreateAccount)
	mux.HandleFunc("/account/login", handler.Login)

	return mux
}
