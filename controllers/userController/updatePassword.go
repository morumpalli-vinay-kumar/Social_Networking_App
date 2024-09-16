package servicecontroller

import (
	"app/middleware/serializers"
	"app/middleware/validators"
	userservice "app/services/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePassword(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var input serializers.PasswordUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validators.ValidatePassword(input.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userservice.UpdatePasswordService(c, input, userID)
}
