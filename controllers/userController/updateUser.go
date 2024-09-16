package servicecontroller

import (
	"app/database"
	"app/middleware/validators"
	"app/models"
	"app/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var updateInput serializers.UserUpdateInput

	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := validators.ValidationCheckUpdate(updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	var office models.OfficeDetails
	var residential models.ResidentialDetails

	if err := database.GORM_DB.Preload("OfficeDetails").Preload("ResidentialDetails").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}

	office = user.OfficeDetails
	residential = user.ResidentialDetails

	response := serializers.BuildUpdateResponse(user, residential, office)

	c.JSON(http.StatusCreated, response)
}
