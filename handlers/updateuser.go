package handlers

import (
	"app/database"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	_, err := utils.ValidateJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.GORM_DB.Where("user_id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Gender = input.Gender
	user.DOB = input.DOB
	user.PersonalEmail = input.PersonalEmail
	user.Password = input.Password
	user.PhoneNumber = input.PhoneNumber
	user.City = input.City
	user.State = input.State
	user.Pincode = input.Pincode
	user.Country = input.Country
	user.IsActive = input.IsActive

	if err := database.GORM_DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
