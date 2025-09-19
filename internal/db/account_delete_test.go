package db

import (
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/stretchr/testify/assert"
)

// Delete Account By ID (Account Exists | Valid)
func TestDABI(t *testing.T) {
	t.Logf("%s: Delete Account By ID", t.Name())
	var account types.Account = types.Account{
		Username: "TestDABI",
		Password: "123",
		Email:    "test@DABI.com",
	}

	id, err := CreateAccount(account)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedId, err := DeleteAccountById(id)

	assert.NoError(t, err)
	assert.Equal(t, id, returnedId)

	assertAccountDoesNotExistId(t, id)

}

// Delete Account By Username (Account Exists | Valid)
func TestDABU(t *testing.T) {
	t.Logf("%s: Delete Account By Username", t.Name())
	var account types.Account = types.Account{
		Username: "TestDABU",
		Password: "123",
		Email:    "Test@DABU.com",
	}

	id, err := CreateAccount(account)

	assert.NoError(t, err)
	assertAccountExist(t, id)

	returnedId, err := DeleteAccountByUsername(account.Username)

	assert.NoError(t, err)
	assert.Equal(t, returnedId, id)
	assertAccountDoesNotExistId(t, id)
}
