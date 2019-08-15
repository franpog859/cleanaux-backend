package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/database"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/handlers"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/secret"
	"github.com/gin-gonic/gin"
)

const (
	externalPort = ":8000"
	internalPort = ":8001"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	databaseClient, secretKey, err := initializeService()
	if err != nil {
		log.Printf("Failed to initialize service: %v", err)
		return
	}

	internalRouter := gin.Default()
	internalHandler := handlers.NewInternalHandler(secretKey)
	internalRouter.POST("/authorize", internalHandler.Authorize)

	externalRouter := gin.Default()
	externalHandler := handlers.NewExternalHandler(databaseClient, secretKey)
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

func initializeService() (database.Client, string, error) {
	databaseClient, err := database.NewClient()
	if err != nil {
		return nil, "", fmt.Errorf("failed to create database client: %v", err)
	}

	secretKey, err := secret.Get()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get secret key: %v", err)
	}

	return databaseClient, secretKey, nil
}
