package servicecontroller

import (
	"app/database"
	"app/models"
	"app/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var user models.UserDetails
	var office models.OfficeDetails
	var residential models.ResidentialDetails

	if err := database.GORM_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := database.GORM_DB.Where("user_id = ?", userID).First(&office).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := database.GORM_DB.Where("user_id = ?", userID).First(&residential).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	response := serializers.BuildUpdateResponse(user, residential, office)

	tx := database.GORM_DB.Begin()

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if err := tx.Where("follower = ? OR following = ?", userID, userID).Delete(&models.Follow{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to delete user"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusAccepted, response)
}
