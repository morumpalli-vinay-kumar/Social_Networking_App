package authcontroller

import (
	"app/database"
	"app/middleware/validators"
	"app/models"
	"app/serializers"
	"app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginInput serializers.Logininput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidationChecklogin(loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundUser models.UserDetails
	if err := database.GORM_DB.Where("email = ?", loginInput.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := foundUser.CheckPassword(loginInput.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(foundUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	expiryTime := time.Now().Add(4 * time.Hour).Format(time.RFC3339)

	userdata := serializers.Loginoutput(foundUser)

	c.JSON(http.StatusOK, gin.H{
		"token": gin.H{
			"key":         token,
			"expiry_time": expiryTime,
		},
		"user": userdata,
	})
}
