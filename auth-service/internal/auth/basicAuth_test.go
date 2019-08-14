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
		header := fmt.Sprintf("Basic %s", encodedPayload)

		// when
		username, password, err := ExtractCredentialsFromHeader(header)

		// then
		require.NoError(t, err)
		assert.Equal(t, expectedUsername, username)
		assert.Equal(t, expectedPassword, password)
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
}
