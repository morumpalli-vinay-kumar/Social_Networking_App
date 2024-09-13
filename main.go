package main

import (
	"app/database"
	"app/middleware"
	"app/routers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func ValidatePassword(pass string) error {
	const minEntropyBits = 80
	return passwordvalidator.Validate(pass, minEntropyBits)
}

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

	router := gin.Default()

	authGroup := router.Group("/auth")
	routers.AuthRoutes(authGroup)

	protected := router.Group("/users")
	protected.Use(middleware.JWTAuthMiddleware())
	routers.UserRoutes(protected)

	router.Run(":8080")

}
