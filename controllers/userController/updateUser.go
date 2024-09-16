package servicecontroller

import (
	"app/middleware/serializers"
	"app/middleware/validators"
	userservice "app/services/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var updateInput serializers.UserUpdateInput

	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := validators.ValidationCheckUpdate(updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userservice.UpdateUserService(c, updateInput, userID)
}
