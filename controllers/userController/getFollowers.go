package servicecontroller

import (
	"app/database"
	"app/models"
	"app/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFollowers(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User ID not found in context"})
		return
	}

	var follows []models.Follow
	if err := database.GORM_DB.Where("following = ?", userID).Find(&follows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve followers list"})
		return
	}

	var userfollowers []serializers.Followoutput
	for _, follow := range follows {
		var user serializers.Followoutput
		if err := database.GORM_DB.Model(&models.UserDetails{}).Where("id = ?", follow.Follower).First(&user).Error; err == nil {
			userfollowers = append(userfollowers, user)
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"followers": userfollowers})
}
