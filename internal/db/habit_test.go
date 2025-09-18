package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
)

func TestMain(m *testing.M) {
	utils.LoadEnvFile()
	dbConfig, err := LoadDatabaseConfig("tst")

	if err != nil {
		os.Exit(1)
	}

	err = dbConfig.Init()

	ClearAllTables()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	code := m.Run()

	os.Exit(code)
}

func TestGetHabitWithTimeValid(t *testing.T) {

}

func TestCreateHabitWithTimeValid(t *testing.T) {
	/*
		habit := types.Habit{
			Name:        "test",
			Description: "Test Description",
			AccountID:   1,
		}
		print(habit)
	*/
}
