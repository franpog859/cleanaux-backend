package main

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	authURL        = "http://auth-service-internal/authorize"
	authHeaderKey  = "Authorization"
	requestTimeout = 10
)

// AuthMiddleware calls auth-service to authorize user.
func AuthMiddleware(context *gin.Context) {
	authHeader := context.GetHeader(authHeaderKey)
	if authHeader == "" {
		log.Println("No Authorization header provided")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	status, err := authorize(authHeader)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if status != http.StatusOK {
		log.Printf("Status: %d, retrying...", status)

		status, err = authorize(authHeader)
		if err != nil {
			log.Println(err)
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if status != http.StatusOK {
			log.Println("Unauthorized")
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
			authHeaderKey: header,
		},
	)

	return status, err
}

func post(postURL string, keyValuePairs map[string]string) (int, error) {
	form := url.Values{}
	for k, v := range keyValuePairs {
		form.Add(k, v)
	}

	req, _ := http.NewRequest("POST", postURL, bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

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
