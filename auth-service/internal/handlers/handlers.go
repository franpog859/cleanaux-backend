package handlers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	jwtTokenSecret = "lasdlkashdakjshd"
	authHeaderKey  = "Authorization"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type user struct {
	id       int
	username string
	password string
}

var users = []user{
	{
		id:       0,
		username: "user0",
		password: "pass0",
	},
	{
		id:       1,
		username: "user1",
		password: "pass1",
	},
}

func Login(context *gin.Context) {
	authHeader := context.GetHeader(authHeaderKey)

	username, password, err := extractCredentialsFromHeader(authHeader)
	if err != nil {
		log.Printf("Error while extracting credentials from header: %v", err)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// TODO: Use database
	valid, err := areCredentialsValid(username, password)
	if err != nil {
		log.Printf("Error while validating credentials: %v", err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !valid {
		log.Printf("Invalid Basic Auth credentials: %s:%s", username, password)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtToken, err := createJWTToken(username)
	if err != nil {
		log.Printf("Failed to create JWT token: %v", err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := tokenResponse{
		Token: jwtToken,
	}

	context.JSON(http.StatusOK, response)
}

func extractCredentialsFromHeader(basicAuthHeader string) (string, string, error) {
	splittedAuthHeader := strings.SplitN(basicAuthHeader, " ", 2)

	if len(splittedAuthHeader) != 2 || splittedAuthHeader[0] != "Basic" {
		return "", "", fmt.Errorf("invalid Basic Auth header: %s", basicAuthHeader)
	}

	authPayload, _ := base64.StdEncoding.DecodeString(splittedAuthHeader[1])
	basicCredentials := strings.SplitN(string(authPayload), ":", 2)

	if len(basicCredentials) != 2 {
		return "", "", fmt.Errorf("invalid Basic Auth credentials: %s", authPayload)
	}

	username, password := basicCredentials[0], basicCredentials[1]

	return username, password, nil
}

func areCredentialsValid(username, password string) (bool, error) {
	for _, user := range users {
		if username == user.username && password == user.password {
			return true, nil
		}
	}

	return false, nil
}

func createJWTToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	signedToken, err := token.SignedString([]byte(jwtTokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Authorize(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}
