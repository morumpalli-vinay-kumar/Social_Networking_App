package handlers

import (
	"app/database"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllActiveUsers(c *gin.Context) {

	var users []models.User

	if err := database.GORM_DB.Where("is_active = ?", true).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var response []gin.H

	for _, user := range users {
		var office models.Office
		if err := database.GORM_DB.Where("user_id = ?", user.UserID).First(&office).Error; err != nil {
			continue
		}

		response = append(response, gin.H{
			"user":   user,
			"office": office,
		})
	}

	c.JSON(http.StatusOK, gin.H{"Data with office is ": response})
}
