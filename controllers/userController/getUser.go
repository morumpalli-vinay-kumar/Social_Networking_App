package servicecontroller

import (
	"app/database"
	"app/models"
	"app/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var user models.UserDetails
	var office models.OfficeDetails
	var residential models.ResidentialDetails

	if err := database.GORM_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}
	if err := database.GORM_DB.Where("user_id = ?", userID).First(&office).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}
	if err := database.GORM_DB.Where("user_id = ?", userID).First(&residential).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}
	response := serializers.BuildUpdateResponse(user, residential, office)

	c.JSON(http.StatusAccepted, response)
}
