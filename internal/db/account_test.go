package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	err := CreateAccount(Account)

	assert.NoError(t, err)
}
