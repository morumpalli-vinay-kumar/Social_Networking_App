package handlers

import (
	"app/database"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {

	_, err := utils.ValidateJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_ID := user.UserID

	if err := database.GORM_DB.Where("user_id = ?", user_ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusOK, gin.H{"message": "User is not active"})
		return
	}

	var office models.Office

	if err := database.GORM_DB.Where("user_id = ?", user_ID).First(&office).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"user": user, "office": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "office": office})
}
