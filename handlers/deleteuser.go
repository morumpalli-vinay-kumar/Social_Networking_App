package handlers

import (
	"app/database"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	// Fetch the user from the database to ensure it exists
	var user models.UserDetails
	if err := database.GORM_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Permanently delete the user record
	if err := database.GORM_DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to delete user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "User and associated details deleted successfully"})
}

// import (
// 	"app/database"
// 	"app/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func DeleteUser(c *gin.Context) {

// 	var input models.User
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var user models.User
// 	if err := database.GORM_DB.Where("user_id = ?", input.UserID).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	user.IsActive = false
// 	if err := database.GORM_DB.Save(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }
