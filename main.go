package main

import (
	"app/database"
	"app/handlers"

	_ "app/migrations"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/socialnetworkapp?sslmode=disable"

	err := database.ConnectToDatabase(dbURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to database ...")

	// if err := goose.SetDialect("postgres"); err != nil {
	// 	panic(err)
	// }

	// if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
	// 	panic(err)
	// }

	r := gin.Default()

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	r.Run(":8080")

}
