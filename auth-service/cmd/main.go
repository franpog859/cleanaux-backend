package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	externalPort = ":8000"
	internalPort = ":8001"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	internalRouter := gin.Default()
	internalRouter.POST("/authorize", authorize)

	externalRouter := gin.Default()
	externalRouter.POST("/login", login)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		internalRouter.Run(internalPort)
	}()

	go func() {
		externalRouter.Run(externalPort)
	}()

	wg.Wait()
}

func login(context *gin.Context) {
	context.JSON(http.StatusOK, "token")
}

func authorize(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}
