package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth_ExtractTokenFromHeader(t *testing.T) {
	t.Run("should return correct token", func(t *testing.T) {
		// given
		expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTY1ODI2MDA5fQ.UMap1wx_B-xGt5PoEvRsQVgaM0b2qhGpsJexLpymm9M"

		header := fmt.Sprintf("%s %s", jwtBearer, expectedToken)

		// when
		token, err := ExtractTokenFromHeader(header)

		// then
		require.NoError(t, err)
		assert.Equal(t, expectedToken, token)
	})

	t.Run("should return error when header structure is invalid", func(t *testing.T) {
		// given
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTY1ODI2MDA5fQ.UMap1wx_B-xGt5PoEvRsQVgaM0b2qhGpsJexLpymm9M"

		cases := []string{
			fmt.Sprintf("%s%s", jwtBearer, token),
			fmt.Sprintf("%s %s", "WrongBearer", token),
			fmt.Sprintf("%s", token),
		}

		for _, header := range cases {
			// when
			_, err := ExtractTokenFromHeader(header)

			// then
			assert.Error(t, err)
		}
	})
}
