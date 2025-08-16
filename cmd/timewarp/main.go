package main

import (
	"fmt"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/server"
)

func main() {

	err := db.Init()

	fmt.Println(err)

	server.Start()

}
