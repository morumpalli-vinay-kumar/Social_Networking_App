package routers

import (
	usercontroller "app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("", usercontroller.GetAllActiveUsers)
	router.GET("/user", usercontroller.GetUserDetails)
	router.GET("/user/following", usercontroller.GetFollowing)
	router.GET("/user/followers", usercontroller.GetFollowers)
	router.PATCH("/user/update", usercontroller.UpdateUser)
	router.DELETE("/user/delete", usercontroller.DeleteUser)
	router.POST("/user/follow", usercontroller.FollowUser)
	router.POST("/user/unfollow", usercontroller.UnfollowUser)
	router.POST("/user/updatepassword", usercontroller.UpdatePassword)
}
