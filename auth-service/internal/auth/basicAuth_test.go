package auth

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth_ExtractCredentialsFromHeader(t *testing.T) {
	t.Run("should return return correct credentials", func(t *testing.T) {
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
