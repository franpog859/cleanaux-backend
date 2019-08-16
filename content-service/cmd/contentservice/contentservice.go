package main

import (
	"log"

	"github.com/franpog859/cleanaux-backend/content-service/internal/database"
	"github.com/franpog859/cleanaux-backend/content-service/internal/handlers"
	"github.com/franpog859/cleanaux-backend/content-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	externalPort = ":8000"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	databaseClient, err := database.NewClient()
	if err != nil {
		log.Printf("Failed to initialize service: %v", err)
		return
	}
	defer databaseClient.Close()

	externalRouter := gin.Default()
	externalHandler := handlers.NewExternalHandler(databaseClient)
	externalRouter.GET("/content", externalHandler.GetContent)
	externalRouter.PUT("/content", externalHandler.PutContent)
	externalRouter.Use(middleware.Auth)

	externalRouter.Run(externalPort)
}
