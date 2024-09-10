package main

import (
	"app/database"
	_ "app/migrations" // This loads the migrations init functions
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
)

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/socialnetworkapp?sslmode=disable"

	err := database.ConnectToDatabase(dbURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to database ...")

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
		panic(err)
	}

	// if err := goose.Down(database.SQL_DB, "migrations"); err != nil {
	// 	panic(err)
	// }

	r := gin.Default()

	r.GET("api/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "helloworld"})
	})
	r.Run()

}
