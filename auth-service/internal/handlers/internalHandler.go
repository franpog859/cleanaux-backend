package handlers

import (
	"log"
	"net/http"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/auth"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/cache"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
	"github.com/gin-gonic/gin"
)

// InternalHandler interface
type InternalHandler interface {
	Authorize(context *gin.Context)
}

type internalHandler struct {
	kubernetesClient kubernetes.Client
	tokenCache       cache.Cache
}

// NewInternalHandler provides InternalHandler interface
func NewInternalHandler(kubernetesClient kubernetes.Client, tokenCache cache.Cache) InternalHandler {
	return &internalHandler{
		kubernetesClient,
		tokenCache,
	}
}

func (ih *internalHandler) Authorize(context *gin.Context) {
	authHeader := context.GetHeader(model.AuthHeaderKey)

	token, err := auth.ExtractTokenFromHeader(authHeader)
	if err != nil {
		log.Printf("Invalid authorization header: %v", err)
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	valid, err := auth.IsTokenValid(token, ih.tokenCache, ih.kubernetesClient)
	if err != nil {
		log.Printf("Error while validating token: %v", err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !valid {
		log.Printf("Invalid authorization token: %s", token)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Status(http.StatusOK)
}
