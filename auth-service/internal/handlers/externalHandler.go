package handlers

import (
	"log"
	"net/http"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/auth"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/database"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
	"github.com/gin-gonic/gin"
)

// ExternalHandler interface
type ExternalHandler interface {
	Login(context *gin.Context)
}

type externalHandler struct {
	databaseClient database.Client
	secretKey      string
}

// NewExternalHandler provides ExternalHandler interface
func NewExternalHandler(dbClient database.Client, secretKey string) ExternalHandler {
	return &externalHandler{
		databaseClient: dbClient,
		secretKey:      secretKey,
	}
}

func (eh *externalHandler) Login(context *gin.Context) {
	authHeader := context.GetHeader(model.AuthHeaderKey)

	username, password, err := auth.ExtractCredentialsFromHeader(authHeader)
	if err != nil {
		log.Printf("Invalid authorization header: %v", err)
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	valid, err := auth.AreCredentialsValid(username, password, eh.databaseClient)
	if err != nil {
		log.Printf("Error while validating credentials: %v", err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !valid {
		log.Printf("Invalid authorization credentials: %s:%s", username, password)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtToken, err := auth.CreateJWTToken(username, eh.secretKey)
	if err != nil {
		log.Printf("Error while creating JWT token: %v", err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.TokenResponse{
		Token: jwtToken,
	}

	context.JSON(http.StatusOK, response)
}
