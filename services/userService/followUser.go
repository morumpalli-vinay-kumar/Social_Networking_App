package userservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FollowUserService(c *gin.Context, userID any, follow serializers.Following) {
	var followerUser models.UserDetails
	if err := database.GORM_DB.Where("id = ?", userID).First(&followerUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You deleted your account"})
		return
	}

	var followingUser models.UserDetails
	if err := database.GORM_DB.Where("id = ?", follow.Following).First(&followingUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User you trying to follow does not exist"})
		return
	}

	var existingFollow models.Follow
	if err := database.GORM_DB.Where("follower = ? AND following = ?", userID, follow.Following).First(&existingFollow).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are already following this user"})
		return
	}

	data := models.Follow{
		Follower:  userID.(uint),
		Following: follow.Following,
	}

	if err := database.GORM_DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
