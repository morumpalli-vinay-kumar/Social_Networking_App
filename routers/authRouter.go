package routers

import (
	authcontroller "app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", authcontroller.Login)
	router.POST("/signup", authcontroller.Signup)
}
