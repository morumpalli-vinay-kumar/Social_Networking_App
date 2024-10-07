package main

import (
	controller "app/controllers"
	"app/database"
	"app/middleware"
	"app/routers"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
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

	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://gorm.io/docs/logger"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	router.GET("/health", controller.Healthcheck)

	authGroup := router.Group("/auth")
	routers.AuthRoutes(authGroup)

	protected := router.Group("/users")
	protected.Use(middleware.JWTAuthMiddleware())
	routers.UserRoutes(protected)

	router.Run(":8080")

}
