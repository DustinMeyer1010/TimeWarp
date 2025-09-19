package db

import (
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAccountByIdExist(t *testing.T) {
	var account types.Account = types.Account{
		Username: "test_valid_id_to_be_deleted",
		Password: "123",
		Email:    "valid@account1tobedeleted.com",
	}

	id, err := CreateAccount(account)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedId, err := DeleteAccountById(id)

	assert.NoError(t, err)
	assert.Equal(t, id, returnedId)

	assertAccountDoesNotExistId(t, id)

}
