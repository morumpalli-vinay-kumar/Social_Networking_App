package servicecontroller

import (
	"app/database"
	"app/models"
	"app/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFollowing(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var follows []models.Follow
	if err := database.GORM_DB.Where("follower = ?", userID).Find(&follows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve following list"})
		return
	}

	var followingUsers []serializers.Followoutput
	for _, follow := range follows {
		var user models.UserDetails
		if err := database.GORM_DB.Where("id = ?", follow.Following).First(&user).Error; err == nil {
			followingUsers = append(followingUsers, serializers.GetFollowingDetails(user))
		}
	}

	c.JSON(http.StatusOK, gin.H{"following": followingUsers})
}
