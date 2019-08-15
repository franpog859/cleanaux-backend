package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	cacheMocks "github.com/franpog859/cleanaux-backend/auth-service/internal/cache/mocks"
	kubernetesMocks "github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes/mocks"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandlers_Authorize(t *testing.T) {
	t.Run("should authorize correct token", func(t *testing.T) {
		// given
		username := "user1"
		secret := "7LEFxuMcVuFnz8T0ipX6QbJD6xZd7qsp94JCBnVXsdcOUBaMR0hk5Z4bsCvjYHN"

		kubernetesClient := &kubernetesMocks.Client{}
		kubernetesClient.On("GetSecret").Return(secret)

		tokenCache := &cacheMocks.Cache{}
		tokenCache.On("GetSecret").Return(secret)

		router := gin.Default()
		handler := NewInternalHandler(kubernetesClient, tokenCache)
		router.POST("/authorize", handler.Authorize)

		req, _ := http.NewRequest("POST", "/authorize", nil)

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(5 * time.Hour).Unix(),
			},
		}).SignedString([]byte(secret))
		require.NoError(t, err)

		header := fmt.Sprintf("Bearer %s", token)
		req.Header.Add(model.AuthHeaderKey, header)

		expectedCode := http.StatusOK

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		require.Equal(t, expectedCode, resp.Code)
	})

	t.Run("should unauthorize invalid token", func(t *testing.T) {
		// given
		secret := "7LEFxuMcVuFnz8T0ipX6QbJD6xZd7qsp94JCBnVXsdcOUBaMR0hk5Z4bsCvjYHN"

		kubernetesClient := &kubernetesMocks.Client{}
		kubernetesClient.On("GetSecret").Return(secret)

		tokenCache := &cacheMocks.Cache{}
		tokenCache.On("GetSecret").Return(secret)
		tokenCache.On("SetSecret", secret).Return()

		router := gin.Default()
		handler := NewInternalHandler(kubernetesClient, tokenCache)
		router.POST("/authorize", handler.Authorize)

		req, _ := http.NewRequest("POST", "/authorize", nil)

		token := "some.wrong.token"
		header := fmt.Sprintf("Bearer %s", token)
		req.Header.Add(model.AuthHeaderKey, header)

		expectedCode := http.StatusUnauthorized

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		require.Equal(t, expectedCode, resp.Code)
	})
}
