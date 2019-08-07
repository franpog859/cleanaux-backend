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

	internalRouter := gin.Default()
	internalRouter.POST("/authorize", handlers.Authorize)

	externalRouter := gin.Default()
	externalRouter.POST("/login", handlers.Login)

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
