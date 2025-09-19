package db

import (
	"fmt"
	"os"
	"testing"

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

// Takes two accounts and compare them to make sure they are the same account
//
// Parameter:
//   - account1: account found in the database with hashed password
//   - account2: account used to store in database with unhashed password
func assertAccountsAreEqual(t *testing.T, account1 types.Account, account2 types.Account) {
	assert.Equal(t, account1.Email, account2.Email)
	assert.Equal(t, account1.Username, account2.Username)
	fmt.Println(account1.Password, account2.Password)
	assert.True(t, account1.CheckPassword(account2))
}

// Takes in an account and makes sure all the fields are empty/default values
func assertAccountIsEmpty(t *testing.T, account types.Account) {
	assert.Equal(t, account.ID, 0)
	assert.Equal(t, account.Username, "")
	assert.Equal(t, account.Password, "")
	assert.Equal(t, account.Email, "")
	assert.True(t, account.CreationDate.IsZero())
}

// Check to make sure username does not exist inside of database
func assertAccountDoesNotExistUsername(t *testing.T, username string) {
	account, err := GetAccountByUsername(username)
	assertAccountIsEmpty(t, account)
	assert.Error(t, err)
}

func assertAccountDoesNotExistId(t *testing.T, id int) {
	account, err := GetAccountByID(id)
	assertAccountIsEmpty(t, account)
	assert.Error(t, err)
}

// Makes sure the id exist inside of database
func assertAccountExist(t *testing.T, id int) {
	account, err := GetAccountByID(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, account.Username, "")
	assert.NotEqual(t, account.ID, -1)
	assert.NotEmpty(t, account.Password)
	assert.NotEmpty(t, account.Email)
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
