package handlers

import (
	"app/database"
	"app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var user models.UserDetails
	var office models.OfficeDetails
	var residential models.ResidentialDetails

	if err := database.GORM_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		fmt.Println("user table ")
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}
	if err := database.GORM_DB.Where("user_id = ?", userID).First(&office).Error; err != nil {
		fmt.Println("office table ")
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}
	if err := database.GORM_DB.Where("user_id = ?", userID).First(&residential).Error; err != nil {
		fmt.Println("resi table ")
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id":        user.ID,
		"email":          user.Email,
		"last_modified":  user.UpdatedAt.Format(time.RFC3339),
		"first_name":     user.FirstName,
		"last_name":      user.LastName,
		"date_of_birth":  user.DateOfBirth,
		"gender":         user.Gender,
		"marital_status": user.MaritalStatus,

		"residential_details": gin.H{
			"address":      residential.Address,
			"city":         residential.City,
			"state":        residential.State,
			"country":      residential.Country,
			"contact_no_1": residential.ContactNo1,
			"contact_no_2": residential.ContactNo2,
		},

		"office_details": gin.H{
			"employee_code": office.EmployeeCode,
			"address":       office.Address,
			"city":          office.City,
			"state":         office.State,
			"country":       office.Country,
			"contact_no":    office.ContactNo,
			"email":         office.Email,
			"name":          office.Name,
		},
	})
}
