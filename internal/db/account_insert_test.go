package db

import (
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/stretchr/testify/assert"
)

// Create Account (Valid)
func TestCA(t *testing.T) {
	t.Logf("%s: Create Account", t.Name())
	var account models.Account = models.Account{
		Username: "TestCA",
		Password: "123",
		Email:    "test@CA.com",
	}
	id, err := CreateAccount(account)

	assert.NoError(t, err)
	assertAccountExist(t, id)
}

// Create Account Missing Email (invalid)
func TestCAME(t *testing.T) {
	t.Logf("%s: Create Account Missing Email", t.Name())
	var account models.Account = models.Account{
		Username: "TestCAME",
		Password: "123",
		Email:    "",
	}

	id, err := CreateAccount(account)
	assert.Equal(t, id, -1)
	assert.Error(t, err)

	assertAccountDoesNotExistUsername(t, account.Username)

}

// Create Account Invalid Email (invalid)
func TestCAIE(t *testing.T) {
	t.Logf("%s: Create Account Invalid Email", t.Name())
	var account models.Account = models.Account{
		Username: "TestCAIE",
		Password: "123",
		Email:    "test@CAIE",
	}

	id, err := CreateAccount(account)

	assert.Equal(t, id, -1)
	assert.Error(t, err)

	assertAccountDoesNotExistUsername(t, account.Username)

}

// Create Account Same Email
func TestCASE(t *testing.T) {
	t.Logf("%s: Create Account Same Email", t.Name())
	var account models.Account = models.Account{
		Username: "TestCASE",
		Password: "123",
		Email:    "test@CASE.com",
	}

	var account1 models.Account = models.Account{
		Username: "TestCASE1",
		Password: "123",
		Email:    "test@CASE.com",
	}

	id, err := CreateAccount(account)

	assertAccountExist(t, id)
	assert.NoError(t, err)

	id, err = CreateAccount(account1)

	assert.Equal(t, id, -1)
	assert.Error(t, err)
	assertAccountDoesNotExistUsername(t, account1.Username)
}

// Create Two Different Accounts (valid)
func TestCTDA(t *testing.T) {
	t.Logf("%s: Create Two Different Accounts", t.Name())
	var account models.Account = models.Account{
		Username:     "TestCTDA",
		Password:     "123",
		Email:        "test@CTDA.com",
		RefreshToken: "",
	}

	var account1 models.Account = models.Account{
		Username:     "TestCTDA1",
		Password:     "123",
		Email:        "test@CTDA1.com",
		RefreshToken: "",
	}

	id, err := CreateAccount(account)

	assertAccountExist(t, id)
	assert.NoError(t, err)

	id, err = CreateAccount(account1)

	assertAccountExist(t, id)
	assert.NoError(t, err)
}

// Create Account Missing Username (invalid)
func TestCAMU(t *testing.T) {
	t.Logf("%s: Create Account Missing Username", t.Name())
	var account models.Account = models.Account{
		Username:     "",
		Password:     "123",
		Email:        "test@CAMU.com",
		RefreshToken: "",
	}

	id, err := CreateAccount(account)

	assert.Error(t, err)
	assert.Equal(t, id, -1)

}
