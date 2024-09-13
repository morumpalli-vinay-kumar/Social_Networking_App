package servicecontroller

import (
	"app/database"
	"app/models"
	"app/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	var follow serializers.Following
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	if err := c.ShouldBindJSON(&follow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var followerUser models.UserDetails
	if err := database.GORM_DB.First(&followerUser, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You deleted your account"})
		return
	}

	var followingUser models.UserDetails
	if err := database.GORM_DB.First(&followingUser, follow.Following).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The user you are trying to follow does not exist"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create follow relationship"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Follow relationship created"})
}
