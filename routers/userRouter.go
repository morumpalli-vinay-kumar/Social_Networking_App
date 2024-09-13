package routers

import (
	usercontroller "app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("", usercontroller.GetAllActiveUsers)
	router.GET("/user", usercontroller.GetUserDetails)
	router.GET("/user/following", usercontroller.GetFollowing)
	router.GET("/user/followers", usercontroller.GetFollowers)
	router.PATCH("/user", usercontroller.UpdateUser)
	router.DELETE("/user", usercontroller.DeleteUser)

	router.POST("/user/unfollow", usercontroller.UnfollowUser)
	router.POST("/user/updatepassword", usercontroller.UpdatePassword)
}
