package handlers

import (
	"app/database"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/nyaruka/phonenumbers"
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

	if !govalidator.IsEmail(request.User.PersonalEmail) {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Email ": request.User.PersonalEmail})
		return
	}

	num, err := phonenumbers.Parse(request.User.PhoneNumber, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Phone Number ": request.User.PhoneNumber})
		return
	}

	if !phonenumbers.IsValidNumber(num) {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Phone Number ": num})
		return
	}

	var existingUser models.User
	if err := database.GORM_DB.Where("user_id = ?", request.User.UserID).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists!"})
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
	errr := database.GORM_DB.Where("user_id = ?", request.Office.UserID).First(&existingOffice).Error
	if errr == nil {
		if err := database.GORM_DB.Model(&existingOffice).Updates(request.Office).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists!"})
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
