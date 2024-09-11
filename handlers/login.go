package handlers

import (
	"app/database"
	"app/models"
	"app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var input models.User

	_, err := utils.ValidateJWTFromHeader(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.GORM_DB.Where("user_id= ?", input.UserID).First(&user).Error; err != nil {
		// fmt.Println("error 1")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user id or password"})
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		fmt.Println("error 2")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "Data": user})
}
