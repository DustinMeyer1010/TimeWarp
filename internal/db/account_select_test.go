package db

import (
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountByIDExist(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_valid_id",
		Password: "123",
		Email:    "valid@account1.com",
	}

	id, err := CreateAccount(account)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedAccount, err := GetAccountByID(id)

	assert.NoError(t, err)
	assertAccountsAreEqual(t, returnedAccount, account)

}

func TestGetAccountByIDNotExist(t *testing.T) {
	DeleteAccountById(1)

	returnedAccount, err := GetAccountByID(1)

	assert.Error(t, err)
	assertAccountIsEmpty(t, returnedAccount)
}
