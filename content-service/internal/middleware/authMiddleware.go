package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/franpog859/cleanaux-backend/content-service/internal/model"
	"github.com/gin-gonic/gin"
)

const (
	authURL        = "http://auth-service-internal/authorize"
	requestTimeout = 10
)

// Auth calls auth-service to authorize user
func Auth(context *gin.Context) {
	authHeader := context.GetHeader(model.AuthHeaderKey)
	if authHeader == "" {
		log.Printf("No Authorization header provided")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	status, err := authorize(authHeader)
	if err != nil {
		log.Printf("Error while authorizing user: %v", err)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if status != http.StatusOK {
		log.Printf("Status: %d, retrying...", status)

		status, err = authorize(authHeader)
		if err != nil {
			log.Printf("Error while authorizing user: %v", err)
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if status != http.StatusOK {
			log.Printf("User is not authorized. Status: %d", status)
			context.AbortWithStatus(status)
			return
		}
	}

	context.Next()
}

func authorize(header string) (int, error) {
	status, err := post(
		authURL,
		map[string]string{
			model.AuthHeaderKey: header,
		},
	)

	return status, err
}

func post(postURL string, headers map[string]string) (int, error) {
	req, _ := http.NewRequest("POST", postURL, nil)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * requestTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
