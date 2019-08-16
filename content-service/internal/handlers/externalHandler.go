package handlers

import (
	"log"
	"net/http"

	"github.com/franpog859/cleanaux-backend/content-service/internal/convert"
	"github.com/franpog859/cleanaux-backend/content-service/internal/database"
	"github.com/franpog859/cleanaux-backend/content-service/internal/model"
	"github.com/gin-gonic/gin"
)

// ExternalHandler interface
type ExternalHandler interface {
	GetContent(*gin.Context)
	PutContent(*gin.Context)
}

type externalHandler struct {
	databaseClient database.Client
}

// NewExternalHandler provides ExternalHandler interface
func NewExternalHandler(dbClient database.Client) ExternalHandler {
	return &externalHandler{
		databaseClient: dbClient,
	}
}

func (eh *externalHandler) GetContent(context *gin.Context) {
	items, err := eh.databaseClient.GetAllItems()
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	content, err := convert.ContentFromItems(items)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, content)
}

func (eh *externalHandler) PutContent(context *gin.Context) {
	var requestBody model.ContentRequest
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updateItemInput := convert.UpdateItemFromContentRequest(requestBody)

	err = eh.databaseClient.UpdateItem(updateItemInput)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusNoContent)
}
