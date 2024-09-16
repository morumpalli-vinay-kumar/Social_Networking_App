package userservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePasswordService(c *gin.Context, input serializers.PasswordUpdateInput, userID any) {

	type UserFields struct {
		Password string `gorm:"column:password"`
	}

	var userFields UserFields
	if err := database.GORM_DB.Model(&models.UserDetails{}).Select("password").Where("id = ?", userID).First(&userFields).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userpassword := userFields.Password

	if err := bcrypt.CompareHashAndPassword([]byte(userpassword), []byte(input.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	if err := database.GORM_DB.Model(&models.UserDetails{}).Where("id = ?", userID).Update("password", string(hashedPassword)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Password updated successfully"})
}
