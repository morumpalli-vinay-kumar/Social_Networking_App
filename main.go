package main

import (
	authcontroller "app/controllers/authController"
	servicecontroller "app/controllers/serviceController"
	"app/database"
	"app/middleware"
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

	r := gin.Default()

	r.GET("/health", servicecontroller.Healthcheck)

	r.POST("/signup", authcontroller.Signup)
	r.POST("/login", authcontroller.Login)

	protected := r.Group("/")

	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/users", servicecontroller.GetAllActiveUsers)
	protected.GET("/user", servicecontroller.GetUserDetails)
	protected.GET("/user/following", servicecontroller.GetFollowing)

	protected.PATCH("/user", servicecontroller.UpdateUser)
	protected.DELETE("/user", servicecontroller.DeleteUser)

	protected.POST("/user/follow", servicecontroller.FollowUser)
	protected.POST("/user/unfollow", servicecontroller.UnfollowUser)
	protected.POST("/user/updatepassword", servicecontroller.UpdatePassword)

	r.Run(":8080")

}
