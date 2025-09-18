package db

import (
	"fmt"
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/stretchr/testify/assert"
)

// Check for a valid account creation with no alteration to any fields
func TestCreateAccount(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_account",
		Password: "123",
		Email:    "test@test.com",
	}
	id, err := CreateAccount(account)

	assert.NoError(t, err)
	assertAccountExist(t, id)
}

func TestCreateAccountMissingEmail(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_missing_email",
		Password: "123",
		Email:    "",
	}

	id, err := CreateAccount(account)
	assert.Equal(t, id, -1)
	assert.Error(t, err)

	assertAccountDoesNotExist(t, account.Username)

}

func TestCreateAccountInvalidEmail(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_account",
		Password: "123",
		Email:    "local@test",
	}

	id, err := CreateAccount(account)

	assert.Equal(t, id, -1)
	assert.Error(t, err)

	assertAccountDoesNotExist(t, account.Username)

}

func TestCreateAccountSameEmail(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_account1",
		Password: "123",
		Email:    "local@test.com",
	}

	var account1 types.Account = types.Account{
		Username: "test_account",
		Password: "123",
		Email:    "local@test.com",
	}

	id, err := CreateAccount(account)

	assertAccountExist(t, id)
	assert.NoError(t, err)

	id, err = CreateAccount(account1)

	assert.Equal(t, id, -1)
	assert.NoError(t, err)
	assertAccountDoesNotExist(t, account1.Username)
}

func TestCreateTwoDifferentAccounts(t *testing.T) {
	var account types.Account = types.Account{
		Username:     "test_account",
		Password:     "123",
		Email:        "local@test.com",
		RefreshToken: "",
	}

	var account1 types.Account = types.Account{
		Username:     "test1_account",
		Password:     "123",
		Email:        "local@test1.com",
		RefreshToken: "",
	}

	id, err := CreateAccount(account)

	assertAccountExist(t, id)
	assert.NoError(t, err)

	id, err = CreateAccount(account1)

	assertAccountExist(t, id)
	assert.NoError(t, err)
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
func assertAccountDoesNotExist(t *testing.T, username string) {
	account, err := GetAccountByUsername(username)
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
