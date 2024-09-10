package handlers

import (
	"app/database"
	"app/models"
	"app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var request struct {
		User   models.User   `json:"user"`
		Office models.Office `json:"office"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while parsing office": err.Error()})
		return
	}

	if err := request.User.HashPassword(request.User.Password); err != nil {
		fmt.Println("error 1")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := database.GORM_DB.Create(&request.User).Error; err != nil {
		fmt.Println("error 2")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	if err := database.GORM_DB.Create(&request.Office).Error; err != nil {
		fmt.Println("error 3")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create office: " + err.Error()})
		return
	}

	token, err := utils.GenerateJWT(request.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.GORM_DB.Where("user_id= ?", input.UserID).First(&user).Error; err != nil {
		fmt.Println("error 1")
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
