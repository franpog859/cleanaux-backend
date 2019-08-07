package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InternalHandler interface
type InternalHandler interface {
	Authorize(context *gin.Context)
}

type internalHandler struct {
	kubernetesClient string
}

// NewInternalHandler provides InternalHandler interface
func NewInternalHandler(kubernetesClient string) InternalHandler {
	return &internalHandler{
		kubernetesClient,
	}
}

func (ih *internalHandler) Authorize(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}
