package servicecontroller

import (
	"app/middleware/serializers"
	userservice "app/services/userService"

	"github.com/gin-gonic/gin"
)

func GetAllActiveUsers(c *gin.Context) {

	var allusers []serializers.AllUsers
	userservice.GetAllActiveUsersService(c, allusers)

}
