package db

import (
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/stretchr/testify/assert"
)

// Get Account By Id Exists
func TestGABIE(t *testing.T) {
	t.Logf("%s: Get Account By Id Exists", t.Name())
	var account models.Account = models.Account{
		Username: "TestGABIE",
		Password: "123",
		Email:    "test@GABIE.com",
	}

	id, err := CreateAccount(account)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedAccount, err := GetAccountByID(id)

	assert.NoError(t, err)
	assertAccountsAreEqual(t, returnedAccount, account)

}

// Get Account By Id Not Exist
func TestGABINE(t *testing.T) {
	t.Logf("%s: Get Account By Id Not Exist", t.Name())
	DeleteAccountById(1)

	returnedAccount, err := GetAccountByID(1)

	assert.Error(t, err)
	assertAccountIsEmpty(t, returnedAccount)
}
