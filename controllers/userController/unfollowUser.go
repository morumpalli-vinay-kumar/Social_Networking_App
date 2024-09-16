package servicecontroller

import (
	"app/middleware/serializers"
	userservice "app/services/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnfollowUser(c *gin.Context) {
	var unfollow serializers.Following

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	if err := c.ShouldBindJSON(&unfollow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userservice.UnfollowUserService(c, unfollow, userID)
}
