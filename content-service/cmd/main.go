package main

import (
	"encoding/json"
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

	response, err := json.Marshal(items) // TODO: Modify the data for client. I mean intervalDays to status for example.
	context.JSON(http.StatusOK, gin.H{"data": string(response)})
}

// TODO: Actually everything.
