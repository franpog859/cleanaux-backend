package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	externalPort = ":8000"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	userRouter := router.Group("/user")
	{
		userRouter.GET("/content", userGetContent)
		userRouter.PUT("/content", userPutContent)
	}

	router.Run(externalPort)
}

func userGetContent(context *gin.Context) {
	database := NewDatabaseService()

	items, err := database.GetAllItems()
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	content, err := CreateContentFromItems(items)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, content)
}

func userPutContent(context *gin.Context) {
	var userRequestBody userContentRequest
	err := context.ShouldBindJSON(&userRequestBody)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updateItemInput := CreateUpdateItemInput(userRequestBody)

	database := NewDatabaseService()

	err = database.UpdateItem(updateItemInput)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}
