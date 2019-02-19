// +build integration

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthService(t *testing.T) {
	auth := authService{Base: "http://localhost:8001"}
	// TODO: Add user1:pass1 to the base

	t.Run("should not be able to login with wrong username/password", func(t *testing.T) {
		// given
		username := "user1"
		password := "wrongpassword"

		expectedToken := ""

		// when
		token := auth.Login(username, password).Token

		// then
		assert.Equal(t, expectedToken, token)
	})

	t.Run("should be able to login with correct username/password", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		notExpectedToken := ""

		// when
		token := auth.Login(username, password).Token

		// then
		assert.NotEqual(t, notExpectedToken, token)
	})

	t.Run("should fail authentication for invalid user", func(t *testing.T) {
		// given
		username := "user1"
		password := "wrongpassword"

		// when
		loginResponse := auth.Login(username, password)
		isAuthenticated := auth.Authenticate(username, loginResponse.Token)

		// then
		assert.Equal(t, false, isAuthenticated)
	})

	t.Run("should authenticate valid user", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		// when
		loginResponse := auth.Login(username, password)
		isAuthenticated := auth.Authenticate(username, loginResponse.Token)

		// then
		assert.Equal(t, true, isAuthenticated)
	})

	t.Run("should logout valid user", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		// when
		loginResponse := auth.Login(username, password)
		isAbleToLogout := auth.Logout(username, loginResponse.Token)

		// then
		assert.Equal(t, true, isAbleToLogout)
	})

	t.Run("should fail to authenticate after logout", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		// when
		loginResponse := auth.Login(username, password)
		isAuthenticated := auth.Authenticate(username, loginResponse.Token)
		require.Equal(t, true, isAuthenticated)

		isAbleToLogout := auth.Logout(username, loginResponse.Token)
		require.Equal(t, true, isAbleToLogout)

		isAuthenticated = auth.Authenticate(username, loginResponse.Token)

		// then
		assert.Equal(t, false, isAuthenticated)
	})
}
