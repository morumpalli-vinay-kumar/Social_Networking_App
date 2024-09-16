package userservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnfollowUserService(c *gin.Context, unfollow serializers.Following, userID any) {
	var follow models.Follow
	if err := database.GORM_DB.Where("follower = ? AND following = ?", userID, unfollow.Following).First(&follow).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not following this user"})
		return
	}

	if err := database.GORM_DB.Delete(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Successfully unfollowed the user"})
}
