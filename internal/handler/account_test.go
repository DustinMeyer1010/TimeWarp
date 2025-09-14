package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/DustinMeyer1010/TimeWarp/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	utils.LoadEnvFile()
	dbConfig, err := db.LoadDatabaseConfig("tst")

	if err != nil {
		panic(err.Error())
	}

	err = dbConfig.Init()

	if err != nil {
		panic("unable to load database")
	}

	code := m.Run()

	os.Exit(code)
}

func TestCreateAccount_Success(t *testing.T) {
	account := types.Account{
		Username: "testUser",
		Password: "123",
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/create/account",
		createJSONBody(account),
	)

	w := httptest.NewRecorder()
	CreateAccount(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	acc, err := db.GetAccountByUsername(account.Username)

	assert.NoError(t, err, "error when trying to find account after status ok")
	assert.Equal(t, acc.Username, account.Username)
	assert.True(t, acc.CheckPassword(account))
}

func TestCreateAccount_BadBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/create/account", bytes.NewBufferString("invalid"))
	w := httptest.NewRecorder()

	CreateAccount(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "unable to parse body")
}
