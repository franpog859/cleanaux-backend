package handlers

import (
	"net/http"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/cache"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes"

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
	context.JSON(http.StatusOK, nil)
}
