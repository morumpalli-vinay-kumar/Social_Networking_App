package servicecontroller

import (
	userservice "app/services/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFollowers(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User ID not found in context"})
		return
	}

	userservice.GetFollowersService(c, userID)
}
