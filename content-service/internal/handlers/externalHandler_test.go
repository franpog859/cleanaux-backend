package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	databaseMocks "github.com/franpog859/cleanaux-backend/content-service/internal/database/mocks"
	"github.com/franpog859/cleanaux-backend/content-service/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_GetContent(t *testing.T) {
	t.Run("should return content", func(t *testing.T) {
		// given
		databaseClient := &databaseMocks.Client{}
		databaseClient.On("GetAllItems").Return([]model.Item{
			{
				ID:            1,
				Name:          "name",
				LastUsageDate: time.Now().Format(model.DateLayout),
				IntervalDays:  20,
			},
		}, nil)

		router := gin.Default()
		handler := NewExternalHandler(databaseClient)
		router.GET("/content", handler.GetContent)

		req, _ := http.NewRequest("GET", "/content", nil)

		expectedCode := http.StatusOK
		expectedBody := []model.ContentResponse{}

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		require.Equal(t, expectedCode, resp.Code)
		assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &expectedBody))
	})

	t.Run("should return InternalServerError if failed to get data from database", func(t *testing.T) {
		// given
		databaseClient := &databaseMocks.Client{}
		databaseClient.On("GetAllItems").Return([]model.Item{}, fmt.Errorf("some error"))

		router := gin.Default()
		handler := NewExternalHandler(databaseClient)
		router.GET("/content", handler.GetContent)

		req, _ := http.NewRequest("GET", "/content", nil)

		expectedCode := http.StatusInternalServerError

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		assert.Equal(t, expectedCode, resp.Code)
	})

	t.Run("should return InternalServerError if data from database is incorrect", func(t *testing.T) {
		// given
		databaseClient := &databaseMocks.Client{}
		databaseClient.On("GetAllItems").Return([]model.Item{
			{
				ID:            1,
				Name:          "name",
				LastUsageDate: time.Now().Format(model.DateLayout),
				IntervalDays:  -1,
			},
		}, nil)

		router := gin.Default()
		handler := NewExternalHandler(databaseClient)
		router.GET("/content", handler.GetContent)

		req, _ := http.NewRequest("GET", "/content", nil)

		expectedCode := http.StatusInternalServerError

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		assert.Equal(t, expectedCode, resp.Code)
	})
}

func TestHandlers_PutContent(t *testing.T) {
	t.Run("should return NoContent if everything is fine", func(t *testing.T) {
		// given
		updateItem := model.UpdateItem{
			ID:            1,
			LastUsageDate: time.Now().Format(model.DateLayout),
		}

		databaseClient := &databaseMocks.Client{}
		databaseClient.On("UpdateItem", updateItem).Return(nil)

		router := gin.Default()
		handler := NewExternalHandler(databaseClient)
		router.PUT("/content", handler.PutContent)

		payload := model.ContentRequest{
			ID: 1,
		}
		requestBody := &bytes.Buffer{}
		json.NewEncoder(requestBody).Encode(payload)

		req, _ := http.NewRequest("PUT", "/content", requestBody)

		expectedCode := http.StatusNoContent

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		require.Equal(t, expectedCode, resp.Code)
	})

	t.Run("should return BadRequest if no body provided", func(t *testing.T) {
		// given
		updateItem := model.UpdateItem{}

		databaseClient := &databaseMocks.Client{}
		databaseClient.On("UpdateItem", updateItem).Return(nil)

		router := gin.Default()
		handler := NewExternalHandler(databaseClient)
		router.PUT("/content", handler.PutContent)

		req, _ := http.NewRequest("PUT", "/content", nil)

		expectedCode := http.StatusBadRequest

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		require.Equal(t, expectedCode, resp.Code)
	})

	t.Run("should return InternalServerError if error in database occurred", func(t *testing.T) {
		// given
		updateItem := model.UpdateItem{
			ID:            1,
			LastUsageDate: time.Now().Format(model.DateLayout),
		}

		databaseClient := &databaseMocks.Client{}
		databaseClient.On("UpdateItem", updateItem).Return(fmt.Errorf("some error"))

		router := gin.Default()
		handler := NewExternalHandler(databaseClient)
		router.PUT("/content", handler.PutContent)

		payload := model.ContentRequest{
			ID: 1,
		}
		requestBody := &bytes.Buffer{}
		json.NewEncoder(requestBody).Encode(payload)

		req, _ := http.NewRequest("PUT", "/content", requestBody)

		expectedCode := http.StatusInternalServerError

		// when
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// then
		require.Equal(t, expectedCode, resp.Code)
	})
}
