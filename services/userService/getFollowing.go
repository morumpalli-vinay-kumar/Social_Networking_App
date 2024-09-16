package userservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFollowingService(c *gin.Context, userID any) {
	var follows []models.Follow
	if err := database.GORM_DB.Where("follower = ?", userID).Find(&follows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve following list"})
		return
	}

	var followingUsers []serializers.Followoutput
	for _, follow := range follows {
		var user serializers.Followoutput
		if err := database.GORM_DB.Model(models.UserDetails{}).Where("id = ?", follow.Following).First(&user).Error; err == nil {
			followingUsers = append(followingUsers, user)
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"following": followingUsers})
}
