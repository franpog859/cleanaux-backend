package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/auth"
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

// ExternalHandler interface
type ExternalHandler interface {
	Login(context *gin.Context)
}

type externalHandler struct {
	databaseClient   string
	kubernetesClient string
}

// NewExternalHandler provides ExternalHandler interface
func NewExternalHandler(databaseClient string, kubernetesClient string) ExternalHandler {
	return &externalHandler{
		databaseClient,
		kubernetesClient,
	}
}

func (eh *externalHandler) Login(context *gin.Context) {
	authHeader := context.GetHeader(authHeaderKey)

	username, password, err := auth.ExtractCredentialsFromHeader(authHeader)
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
