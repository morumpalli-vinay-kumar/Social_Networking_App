package servicecontroller

import (
	"app/database"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllActiveUsers(c *gin.Context) {

	var allusers []struct {
		ID    uint   `json:"user_id"`
		Email string `json:"email"`
	}

	if err := database.GORM_DB.Model(&models.UserDetails{}).
		Select("id, email").
		Find(&allusers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Users": allusers})
}
