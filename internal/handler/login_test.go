package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin_Success(t *testing.T) {
	account := models.Account{
		Username: "testUser",
		Password: "123",
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/account/login",
		createJSONBody(account),
	)
	w := httptest.NewRecorder()

	Login(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Header().Get("Authorization"))

	cookies := w.Result().Cookies()
	found := false
	for _, c := range cookies {
		if c.Name == "refresh_token" {
			found = true
			assert.NotEmpty(t, c.Value)
			assert.True(t, c.HttpOnly)
			assert.True(t, c.Secure)
			assert.Equal(t, "/", c.Path)
			assert.Equal(t, http.SameSiteStrictMode, c.SameSite)
			assert.WithinDuration(t, time.Now().Add(7*24*time.Hour), c.Expires, time.Hour)
		}
	}
	assert.True(t, found, "refresh_token cookie should be set")
}

func TestLogin_BadBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/account/login", bytes.NewBufferString("invalid"))
	w := httptest.NewRecorder()

	Login(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "unable to parse body")
}

func TestLogin_AccountNotFound(t *testing.T) {
	account := models.Account{
		Username: "NoFound",
		Password: "123",
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/account/login",
		createJSONBody(account))
	w := httptest.NewRecorder()

	Login(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Failed to add to db")
}

func TestLogin_InvalidPassword(t *testing.T) {
	account := models.Account{
		Username: "testUser",
		Password: "1234",
	}
	req := httptest.NewRequest(
		http.MethodPost,
		"/account/login",
		createJSONBody(account),
	)
	w := httptest.NewRecorder()

	Login(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid password")
}

func createJSONBody(account models.Account) *bytes.Buffer {
	body, _ := json.Marshal(account)
	return bytes.NewBuffer(body)
}
