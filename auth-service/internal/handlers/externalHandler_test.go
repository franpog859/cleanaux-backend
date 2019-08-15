package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	databaseMocks "github.com/franpog859/cleanaux-backend/auth-service/internal/database/mocks"
	kubernetesMocks "github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes/mocks"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_Login(t *testing.T) {
	t.Run("should return token", func(t *testing.T) {
		// given
		username := "user1"
		password := "pass1"

		databaseClient := &databaseMocks.Client{}
		databaseClient.On("GetAuthorizedUsers", username, password).Return([]model.User{
			{
				Username: username,
				Password: password,
			},
		}, nil)

		secret := "7LEFxuMcVuFnz8T0ipX6QbJD6xZd7qsp94JCBnVXsdcOUBaMR0hk5Z4bsCvjYHN"

		kubernetesClient := &kubernetesMocks.Client{}
		kubernetesClient.On("GetSecret").Return(secret)

		router := gin.Default()
		handler := NewExternalHandler(databaseClient, kubernetesClient)
		router.POST("/login", handler.Login)

		req, _ := http.NewRequest("POST", "/login", nil)

		payload := fmt.Sprintf("%s:%s", username, password)
		encodedPayload := base64.StdEncoding.EncodeToString([]byte(payload))
		header := fmt.Sprintf("Basic %s", encodedPayload)
		req.Header.Add(model.AuthHeaderKey, header)

		expectedCode := http.StatusOK
		expectedBody := model.TokenResponse{}

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		require.Equal(t, expectedCode, resp.Code)
		assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &expectedBody))
	})
}
