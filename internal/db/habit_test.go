package db

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
	"github.com/stretchr/testify/assert"
)

var Account1 types.Account = types.Account{
	Username: "Habit_Account1",
	Password: "123",
	Email:    "HabitAccount1@test.com",
}

var Account2 types.Account = types.Account{
	Username: "Habit_Account2",
	Password: "123",
	Email:    "HabitAccount2@test.com",
}

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
	// Accounts that are created for test purposes
	Account1.ID, err = CreateAccount(Account1)

	if err != nil {
		fmt.Println(err)
		panic("unable to test without account creation")
	}
	Account2.ID, err = CreateAccount(Account2)

	if err != nil {
		panic("unable to test without account creation")
	}

	fmt.Println(Account1.ID)

	code := m.Run()

	os.Exit(code)
}

func TestCreateHabitWithTimeValid(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)
}

func TestCreateSameHabitTwiceValid(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "TwoHabits",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}
	id, err := CreateHabitWithTime(habit)
	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	id2, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	assert.NotEqual(t, id, id2)
}

func TestCreateHabitWithoutTime(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:        "TwoHabits",
		Description: "This is a valid habit",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)
}

func TestGetHabitWithTime(t *testing.T) {

	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)
	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, habit.AccountID)

	assert.NoError(t, err)

	assertHabitsAreEqual(t, returnedHabit, habit)
}

func TestGetHabitWithTimeNotExist(t *testing.T) {
	DeleteHabitWithTime(1, Account1.ID) // Make sure there is no habit that exist with this habit

	habit, err := GetHabitWithTime(1, Account1.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, habit)

}

func TestGetHabitWithoutTimeNotExist(t *testing.T) {
	DeleteHabitWithoutTime(1, Account1.ID)

	habit, err := GetHabitWithoutTime(1, Account1.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, habit)

}

func TestGetHabitWithTimeExist(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NotEqual(t, id, -1)
	assert.NoError(t, err)

	returnedHabit, err := GetHabitWithTime(id, habit.AccountID)

	assert.NoError(t, err)
	assertHabitsAreEqual(t, returnedHabit, habit)

}

func TestGetHabitWithoutTimeExist(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:        "Test",
		Description: "This is a valid habit",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithoutTime(id, habit.AccountID)

	assert.NoError(t, err)
	assertHabitsAreEqual(t, returnedHabit, habit)

}

func TestGetHabitWithTimeIDAndAccountNoMatch(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, Account2.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, returnedHabit)

}

func TestGetHabitWithoutTimeIDandAccountNoMatch(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:        "Test",
		Description: "This is a valid habit",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, Account2.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, returnedHabit)
}

// Takes two habits and compares them to make sure they equal base on all values
func assertHabitsAreEqual(t *testing.T, habit1, habit2 types.Habit) {
	assert.Equal(t, habit1.AccountID, habit2.AccountID)
	assert.Equal(t, habit1.Name, habit2.Name)
	assert.Equal(t, habit1.Description, habit2.Description)
	assert.Equal(t, habit1.CompletionTime, habit2.CompletionTime)
}

// Takes a habit and makes sure that all values are empty/default
func assertHabitIsEmpty(t *testing.T, habit types.Habit) {
	assert.Empty(t, habit.AccountID)
	assert.Empty(t, habit.Name)
	assert.Empty(t, habit.Description)
	assert.True(t, habit.CompletionTime.IsZero())
}
