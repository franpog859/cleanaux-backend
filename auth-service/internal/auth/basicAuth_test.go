package auth

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/database/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth_ExtractCredentialsFromHeader(t *testing.T) {
	t.Run("should return correct credentials", func(t *testing.T) {
		// given
		expectedUsername := "user1"
		expectedPassword := "pass1"

		payload := fmt.Sprintf("%s:%s", expectedUsername, expectedPassword)
		encodedPayload := base64.StdEncoding.EncodeToString([]byte(payload))
		header := fmt.Sprintf("%s %s", basicBearer, encodedPayload)

		// when
		username, password, err := ExtractCredentialsFromHeader(header)

		// then
		require.NoError(t, err)
		assert.Equal(t, expectedUsername, username)
		assert.Equal(t, expectedPassword, password)
	})

	t.Run("should return error when header structure is invalid", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		payload := fmt.Sprintf("%s:%s", username, password)
		encodedPayload := base64.StdEncoding.EncodeToString([]byte(payload))

		cases := []string{
			fmt.Sprintf("%s%s", basicBearer, encodedPayload),
			fmt.Sprintf("%s %s", "WrongBearer", encodedPayload),
			fmt.Sprintf("%s", encodedPayload),
		}

		for _, header := range cases {
			// when
			_, _, err := ExtractCredentialsFromHeader(header)

			// then
			assert.Error(t, err)
		}
	})

	t.Run("should return error when credentials structure is invalid", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		cases := []string{
			base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s%s", username, password))),
			base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s", password))),
			fmt.Sprintf("%s:%s", username, password),
		}

		for _, encodedPayload := range cases {
			header := fmt.Sprintf("%s %s", basicBearer, encodedPayload)

			// when
			_, _, err := ExtractCredentialsFromHeader(header)

			// then
			assert.Error(t, err)
		}
	})
}

func TestAuth_AreCredentialsValid(t *testing.T) {
	t.Run("should validate correct credentials", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		databaseClient := &mocks.Client{}
		databaseClient.On("GetAuthorizedUsers", username, password).Return([]model.User{
			{
				Username: username,
				Password: password,
			},
		}, nil)

		expectedValid := true

		// when
		valid, err := AreCredentialsValid(username, password, databaseClient)

		// then
		require.NoError(t, err)
		assert.Equal(t, expectedValid, valid)
	})

	t.Run("should validate multiple correct credentials", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		databaseClient := &mocks.Client{}
		databaseClient.On("GetAuthorizedUsers", username, password).Return([]model.User{
			{
				Username: username,
				Password: password,
			},
			{
				Username: username,
				Password: password,
			},
		}, nil)

		expectedValid := true

		// when
		valid, err := AreCredentialsValid(username, password, databaseClient)

		// then
		require.NoError(t, err)
		assert.Equal(t, expectedValid, valid)
	})

	t.Run("should invalidate invalid credentials", func(t *testing.T) {
		// given
		username := "asd!@#"
		password := "@#$qwe"

		databaseClient := &mocks.Client{}
		databaseClient.On("GetAuthorizedUsers", username, password).Return([]model.User{
			{
				Username: username,
				Password: password,
			},
		}, nil)

		expectedValid := false

		// when
		valid, err := AreCredentialsValid(username, password, databaseClient)

		// then
		require.NoError(t, err)
		assert.Equal(t, expectedValid, valid)
	})

	t.Run("should return error if database error occurred", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		databaseClient := &mocks.Client{}
		databaseClient.On("GetAuthorizedUsers", username, password).Return([]model.User{
			{
				Username: username,
				Password: password,
			},
		}, fmt.Errorf("some error"))

		// when
		_, err := AreCredentialsValid(username, password, databaseClient)

		// then
		assert.Error(t, err)
	})

	t.Run("should invalidate incorrect credentials", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		databaseClient := &mocks.Client{}
		databaseClient.On("GetAuthorizedUsers", username, password).Return([]model.User{}, nil)

		expectedValid := false

		// when
		valid, err := AreCredentialsValid(username, password, databaseClient)

		// then
		require.NoError(t, err)
		assert.Equal(t, expectedValid, valid)
	})
}
