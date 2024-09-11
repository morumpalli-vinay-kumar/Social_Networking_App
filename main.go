package main

import (
	"app/database"
	"app/handlers"
	"app/middleware"
	"log"
	"os"

	_ "app/migrations"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err1 := godotenv.Load("/home/ubuntu/app/goose.env")
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, databaseName, sslmode)
	err2 := database.ConnectToDatabase(dbURL)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Connected to Database ----> ", databaseName)

	r := gin.Default()
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/users", handlers.GetAllActiveUsers)
	protected.GET("/users/:user_id", handlers.GetUserDetails)

	protected.PATCH("/users", handlers.UpdateUser)
	protected.DELETE("/users", handlers.DeleteUser)

	r.Run(":8080")

}
