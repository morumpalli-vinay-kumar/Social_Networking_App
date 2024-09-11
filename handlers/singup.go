package handlers

import (
	"app/database"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var request struct {
		User   models.User   `json:"user"`
		Office models.Office `json:"office"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while parsing request": err.Error()})
		return
	}

	if err := request.User.HashPassword(request.User.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := database.GORM_DB.Create(&request.User).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	request.Office.UserID = request.User.UserID

	var existingOffice models.Office
	err := database.GORM_DB.Where("user_id = ?", request.Office.UserID).First(&existingOffice).Error
	if err == nil {
		if err := database.GORM_DB.Model(&existingOffice).Updates(request.Office).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update office: " + err.Error()})
			return
		}
	} else {
		if err := database.GORM_DB.Create(&request.Office).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create office: " + err.Error()})
			return
		}
	}

	token, err := utils.GenerateJWT(request.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
