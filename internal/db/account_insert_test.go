package db

import (
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_valid_account_created",
		Password: "123",
		Email:    "single@account.com",
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

	assertAccountDoesNotExistUsername(t, account.Username)

}

func TestCreateAccountInvalidEmail(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_invalid_email",
		Password: "123",
		Email:    "local@test",
	}

	id, err := CreateAccount(account)

	assert.Equal(t, id, -1)
	assert.Error(t, err)

	assertAccountDoesNotExistUsername(t, account.Username)

}

func TestCreateAccountSameEmail(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_valid_account",
		Password: "123",
		Email:    "same@email.com",
	}

	var account1 types.Account = types.Account{
		Username: "test_invalid_email_exist",
		Password: "123",
		Email:    "same@email.com",
	}

	id, err := CreateAccount(account)

	assertAccountExist(t, id)
	assert.NoError(t, err)

	id, err = CreateAccount(account1)

	assert.Equal(t, id, -1)
	assert.Error(t, err)
	assertAccountDoesNotExistUsername(t, account1.Username)
}

func TestCreateTwoDifferentAccounts(t *testing.T) {
	var account types.Account = types.Account{
		Username:     "test_valid1",
		Password:     "123",
		Email:        "wow@test.com",
		RefreshToken: "",
	}

	var account1 types.Account = types.Account{
		Username:     "test_valid2",
		Password:     "123",
		Email:        "bob@test1.com",
		RefreshToken: "",
	}

	id, err := CreateAccount(account)

	assertAccountExist(t, id)
	assert.NoError(t, err)

	id, err = CreateAccount(account1)

	assertAccountExist(t, id)
	assert.NoError(t, err)
}

func TestCreateAccountMissingUsername(t *testing.T) {
	var account types.Account = types.Account{
		Username:     "",
		Password:     "123",
		Email:        "Missing@Username.com",
		RefreshToken: "",
	}

	id, err := CreateAccount(account)

	assert.Error(t, err)
	assert.Equal(t, id, -1)

}
