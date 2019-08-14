package main

import (
	"sync"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/cache"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/database"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/handlers"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes"
	"github.com/gin-gonic/gin"
)

const (
	externalPort = ":8000"
	internalPort = ":8001"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	kubernetesClient := kubernetes.NewClient()
	databaseClient := database.NewClient()
	tokenCache := cache.New()

	internalRouter := gin.Default()
	internalHandler := handlers.NewInternalHandler(kubernetesClient, tokenCache)
	internalRouter.POST("/authorize", internalHandler.Authorize)

	externalRouter := gin.Default()
	externalHandler := handlers.NewExternalHandler(databaseClient, kubernetesClient)
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
