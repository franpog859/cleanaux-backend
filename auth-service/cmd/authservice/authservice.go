package main

import (
	"fmt"
	"log"
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

	kubernetesClient, databaseClient, tokenCache, err := initializeClients()
	if err != nil {
		log.Printf("Failed to initialize service: %v", err)
		return
	}

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

func initializeClients() (kubernetes.Client, database.Client, cache.Cache, error) {
	kubernetesClient := kubernetes.NewClient()

	databaseClient, err := database.NewClient()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create database client: %v", err)
	}

	tokenCache := cache.New()

	return kubernetesClient, databaseClient, tokenCache, nil
}
