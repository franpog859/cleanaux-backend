package handlers

import (
	"log"
	"net/http"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/database"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/auth"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey = "Authorization"
)

// ExternalHandler interface
type ExternalHandler interface {
	Login(context *gin.Context)
}

type externalHandler struct {
	databaseClient   database.Client
	kubernetesClient kubernetes.Client
}

// NewExternalHandler provides ExternalHandler interface
func NewExternalHandler(databaseClient database.Client, kubernetesClient kubernetes.Client) ExternalHandler {
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

	// TODO: Use database to get users (pass it to the function)
	valid, err := auth.AreCredentialsValid(username, password, eh.databaseClient)
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

	// TODO: Use k8sClient to get jwtTokenSecret
	jwtToken, err := auth.CreateJWTToken(username)
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
