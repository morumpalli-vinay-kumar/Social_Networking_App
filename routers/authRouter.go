package routers

import (
	authcontroller "app/controllers/authController"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.GET("/health", authcontroller.Healthcheck)
	router.POST("/login", authcontroller.Login)
	router.POST("/signup", authcontroller.Signup)
}
