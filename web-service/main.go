package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	auth := authService{Base: "http://localhost:8001"}
	if response := auth.Login(username, password); response.Token != "" {
		context.SetCookie("username", username, 3600, "", "", false, true)
		context.SetCookie("token", response.Token, 3600, "", "", false, true)

		context.JSON(http.StatusOK, response)
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

func logout(context *gin.Context) {
	username, err1 := context.Cookie("username")
	token, err2 := context.Cookie("token")

	auth := authService{Base: "http://localhost:8001"}
	if err1 == nil && err2 == nil && auth.Logout(username, token) {
		context.SetCookie("username", "", -1, "", "", false, true)
		context.SetCookie("token", "", -1, "", "", false, true)

		context.JSON(http.StatusOK, nil)
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

func serveContent(context *gin.Context) {
	username, err1 := context.Cookie("username")
	token, err2 := context.Cookie("token")

	auth := authService{Base: "http://localhost:8001"}
	if err1 == nil && err2 == nil && auth.Authenticate(username, token) {
		context.JSON(http.StatusOK, gin.H{"content": ":)"})
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
