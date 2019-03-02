package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Success!")

}
