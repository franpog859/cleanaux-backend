package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

var users = make(map[string]string)

var seedUsers = []user{ // TODO: Implement some auth here. It is currently NOAuth.
	user{
		Username: "user1",
		Password: "pass1",
	},
	user{
		Username: "user2",
		Password: "pass2",
	},
	user{
		Username: "user3",
		Password: "pass3",
	},
}

const port = ":8000"

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/authenticate", authenticate)
	router.Run(port)
}

func login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	token := validateUser(username, password)

	if token != "" {
		users[username] = token
		context.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

func logout(context *gin.Context) {
	username := context.PostForm("username")
	token := context.PostForm("token")

	value, ok := users[username]
	if ok && value == token {
		delete(users, username)
		context.JSON(http.StatusOK, nil)
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

func authenticate(context *gin.Context) {
	username := context.PostForm("username")
	token := context.PostForm("token")

	value, ok := users[username]
	if ok && value == token {
		context.JSON(http.StatusOK, nil)
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

func validateUser(username, password string) string {
	// TODO: Use database storing users.
	for _, user := range seedUsers {
		if username == user.Username {
			if password == user.Password {
				return generateSessionToken()
			}
			return ""
		}
	}
	return ""
}

func generateSessionToken() string {
	// TODO: Create more secure token
	return strconv.FormatInt(rand.Int63(), 16)
}
