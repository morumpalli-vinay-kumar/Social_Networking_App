package userservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserService(c *gin.Context, updateInput serializers.UserUpdateInput, userID any) {
	updatedFields := map[string]interface{}{
		"first_name":     updateInput.FirstName,
		"last_name":      updateInput.LastName,
		"gender":         updateInput.Gender,
		"date_of_birth":  updateInput.DateOfBirth,
		"marital_status": updateInput.MaritalStatus,
	}

	if err := database.GORM_DB.Model(&models.UserDetails{}).
		Where("id = ?", userID).
		Updates(updatedFields).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	var user models.UserDetails

	if err := database.GORM_DB.Preload("OfficeDetails").Preload("ResidentialDetails").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}

	response := serializers.BuildUpdateResponse(user)

	c.JSON(http.StatusCreated, response)
}
