package handlers

import (
	"app/database"
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
	var user models.UserDetails

	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "only first name, last name, gender,date of birth , Marital Status are allowed to update"})
		return
	}

	if err := database.GORM_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User has been updated already"})
		return
	}

	user.FirstName = updateInput.FirstName
	user.LastName = updateInput.LastName
	user.Gender = updateInput.Gender
	user.DateOfBirth = updateInput.DateOfBirth
	user.MaritalStatus = updateInput.MaritalStatus

	if err := database.GORM_DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to update user"})
		return
	}

	var office models.OfficeDetails
	var residential models.ResidentialDetails

	if err := database.GORM_DB.Where("user_id = ?", userID).First(&office).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": userID})
		return
	}
	if err := database.GORM_DB.Where("user_id = ?", userID).First(&residential).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": userID})
		return
	}

	response := serializers.BuildUpdateResponse(user, residential, office)

	c.JSON(http.StatusCreated, response)
}
