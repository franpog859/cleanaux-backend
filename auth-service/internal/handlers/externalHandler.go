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
func NewExternalHandler(dbClient database.Client, k8sClient kubernetes.Client) ExternalHandler {
	return &externalHandler{
		databaseClient:   dbClient,
		kubernetesClient: k8sClient,
	}
}

func (eh *externalHandler) Login(context *gin.Context) {
	authHeader := context.GetHeader(authHeaderKey)

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

	jwtToken, err := auth.CreateJWTToken(username, eh.kubernetesClient)
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
