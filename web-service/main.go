package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var auth = authService{Base: "localhost:8001"}

const port = ":8000"

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/login", login)
	router.GET("/logout", logout)
	router.GET("/content", serveContent)
	router.Run(port)
}

func login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	response := auth.Login(username, password)
	if response.Token != "" {
		context.SetCookie("username", username, 3600, "", "", false, true)
		context.SetCookie("token", response.Token, 3600, "", "", false, true)

		context.JSON(http.StatusOk, response)
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
