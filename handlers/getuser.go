package handlers

import (
	"app/database"
	"app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {

	userID, check := c.Get("user_id")
	fmt.Println(userID)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}

	var user models.User

	if err := database.GORM_DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusOK, gin.H{"message": "User is not active"})
		return
	}

	var office models.Office

	if err := database.GORM_DB.Where("user_id = ?", userID).First(&office).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"user": user, "office": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "office": office})
}
