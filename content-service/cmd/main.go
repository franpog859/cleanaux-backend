package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8000"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	userRouter := router.Group("/user")
	{
		userRouter.GET("/content", userGetContent)
		//userRouter.PUT("/content", userPutContent)
	}

	router.Run(port)
}

func userGetContent(context *gin.Context) {
	database := NewDatabaseService()

	items, err := database.GetAllItems()
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	content, err := GetContentFromItems(items)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	context.JSON(http.StatusOK, content)
}

// TODO: Actually everything.
