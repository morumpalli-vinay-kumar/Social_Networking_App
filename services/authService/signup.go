package authservice

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignupService(c *gin.Context, req serializers.User) {
	tx := database.GORM_DB.Begin()

	userDetails := models.UserDetails{
		Email:         req.Email,
		Password:      req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		DateOfBirth:   req.DateOfBirth,
		Gender:        req.Gender,
		MaritalStatus: req.MaritalStatus,
	}

	if err := userDetails.HashPassword(userDetails.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := tx.Create(&userDetails).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	residentialDetails := models.ResidentialDetails{
		UserID:     userDetails.ID,
		Address:    req.ResidentialDetails.Address,
		City:       req.ResidentialDetails.City,
		State:      req.ResidentialDetails.State,
		Country:    req.ResidentialDetails.Country,
		ContactNo1: req.ResidentialDetails.ContactNo1,
		ContactNo2: req.ResidentialDetails.ContactNo2,
	}

	if err := tx.Create(&residentialDetails).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	officeDetails := models.OfficeDetails{
		UserID:       userDetails.ID,
		EmployeeCode: req.OfficeDetails.EmployeeCode,
		Address:      req.OfficeDetails.Address,
		City:         req.OfficeDetails.City,
		State:        req.OfficeDetails.State,
		Country:      req.OfficeDetails.Country,
		ContactNo:    req.OfficeDetails.ContactNo,
		Email:        req.OfficeDetails.Email,
		Name:         req.OfficeDetails.Name,
	}

	if err := tx.Create(&officeDetails).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(userDetails.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	expiryTime := time.Now().Add(4 * time.Hour).Format(time.RFC3339)
	tx.Commit()

	response := serializers.BuildUserResponse(userDetails, residentialDetails, officeDetails, token, expiryTime)

	c.JSON(http.StatusCreated, response)
}