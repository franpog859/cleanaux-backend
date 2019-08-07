package main

import (
	"sync"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

const (
	externalPort = ":8000"
	internalPort = ":8001"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// create k8sClient
	// create dbClient

	internalRouter := gin.Default()
	internalHandler := handlers.NewInternalHandler("k8sClient")
	internalRouter.POST("/authorize", internalHandler.Authorize)

	externalRouter := gin.Default()
	externalHandler := handlers.NewExternalHandler("dbClient", "k8sClient")
	externalRouter.POST("/login", externalHandler.Login)

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
