package main

import (
	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/server"
)

func main() {

	db.Init() // handle error later

	server.Start()

}
