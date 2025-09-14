package main

import (
	"fmt"
	"os"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/server"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

func main() {

	utils.LoadEnvFile()
	dbConfig, err := db.LoadDatabaseConfig("dev")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = dbConfig.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	server.Start()
}
