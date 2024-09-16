package servicecontroller

import (
	"app/middleware/serializers"
	userservice "app/services/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	var follow serializers.Following
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User ID not found in context"})
		return
	}

	if err := c.ShouldBindJSON(&follow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userservice.FollowUserService(c, userID, follow)

}
