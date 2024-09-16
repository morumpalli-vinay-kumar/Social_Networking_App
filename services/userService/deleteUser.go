package userservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteService(c *gin.Context, userID any) {
	var user models.UserDetails

	if err := database.GORM_DB.Preload("OfficeDetails").Preload("ResidentialDetails").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.BuildUpdateResponse(user)

	tx := database.GORM_DB.Begin()

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	if err := tx.Delete(&user.ResidentialDetails).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete residential details"})
		return
	}
	if err := tx.Delete(&user.OfficeDetails).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete office details"})
		return
	}

	if err := tx.Where("follower = ? OR following = ?", userID, userID).Delete(&models.Follow{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete follow relations"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusAccepted, response)
}
