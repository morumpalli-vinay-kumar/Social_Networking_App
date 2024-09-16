package userservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllActiveUsersService(c *gin.Context, allusers []serializers.AllUsers) {
	if err := database.GORM_DB.Model(&models.UserDetails{}).
		Select("id, email").
		Find(&allusers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Users": allusers})
}
