package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	port                 = ":8000"
	databaseSourceConfig = "root:password@tcp(mysql:3306)/content"
)

type item struct {
	id            int
	name          string
	lastUsageDate string
	intervalDays  int
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	userRouter := router.Group("/user")
	{
		userRouter.GET("/content", userGetContent)
		//userRouter.PUT("/content", userPutContent)
	}

	router.Run(port)
}

func userGetContent(context *gin.Context) {
	db, err := sql.Open("mysql", databaseSourceConfig)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
	}
	defer db.Close()

	selectDB, err := db.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	response := []item{}
	for selectDB.Next() {
		var id, intervalDays int
		var name, lastUsageDate string
		err := selectDB.Scan(&id, &name, &lastUsageDate, &intervalDays)
		if err != nil {
			fmt.Println(err)
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		itemInstance := item{
			id,
			name,
			lastUsageDate,
			intervalDays,
		}
		response = append(response, itemInstance)
	}

	context.JSON(http.StatusOK, gin.H{"database": response})
}

// TODO: Actually everything.
