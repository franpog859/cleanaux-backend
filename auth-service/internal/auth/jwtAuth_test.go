package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth_CreateJWTToken(t *testing.T) {
	t.Run("should return valid token", func(t *testing.T) {
		// given
		username := "user1"
		jwtKey := "7LEFxuMcVuFnz8T0ipX6QbJD6xZd7qsp94JCBnVXsdcOUBaMR0hk5Z4bsCvjYHN"

		// when
		token, err := CreateJWTToken(username, jwtKey)

		// then
		require.NoError(t, err)

		parsedToken, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		require.NoError(t, err)

		assert.True(t, parsedToken.Valid)
	})
}

func TestAuth_ExtractTokenFromHeader(t *testing.T) {
	t.Run("should return correct token", func(t *testing.T) {
		// given
		expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTY1ODI2MDA5fQ.UMap1wx_B-xGt5PoEvRsQVgaM0b2qhGpsJexLpymm9eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTY1ODI2MDA5fQ.UMap1wx_B-xGt5PoEvRsQVgaM0b2qhGpsJexLpymm9M"

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

func TestAuth_IsTokenValid(t *testing.T) {
	t.Run("should validate valid token", func(t *testing.T) {
		// given
		jwtKey := "7LEFxuMcVuFnz8T0ipX6QbJD6xZd7qsp94JCBnVXsdcOUBaMR0hk5Z4bsCvjYHN"
		username := "user1"

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(5 * time.Hour).Unix(),
			},
		}).SignedString([]byte(jwtKey))
		require.NoError(t, err)

		// when
		valid, err := IsTokenValid(token, jwtKey)

		// then
		require.NoError(t, err)
		assert.True(t, valid)
	})

	t.Run("should invalidate not valid token", func(t *testing.T) {
		// given
		jwtKey := "7LEFxuMcVuFnz8T0ipX6QbJD6xZd7qsp94JCBnVXsdcOUBaMR0hk5Z4bsCvjYHN"
		token := "some.wrong.token"

		// when
		valid, err := IsTokenValid(token, jwtKey)

		// then
		assert.Error(t, err)
		assert.False(t, valid)
	})

	t.Run("should invalidate expired token", func(t *testing.T) {
		// given
		jwtKey := "7LEFxuMcVuFnz8T0ipX6QbJD6xZd7qsp94JCBnVXsdcOUBaMR0hk5Z4bsCvjYHN"
		username := "user1"

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(-5 * time.Hour).Unix(),
			},
		}).SignedString([]byte(jwtKey))
		require.NoError(t, err)

		// when
		valid, err := IsTokenValid(token, jwtKey)

		// then
		assert.Error(t, err)
		assert.False(t, valid)
	})
}
