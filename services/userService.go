package services

import (
	"app/database"
	"app/middleware/serializers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func DeleteService(c *gin.Context, userID any) {
	var user models.UserDetails

	if err := database.GORM_DB.Preload("OfficeDetails").Preload("ResidentialDetails").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.BuildUpdateResponse(user)

	tx := database.GORM_DB.Begin()

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if err := tx.Delete(&user.ResidentialDetails).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete residential details"})
		return
	}
	if err := tx.Delete(&user.OfficeDetails).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete office details"})
		return
	}

	if err := tx.Where("follower = ? OR following = ?", userID, userID).Delete(&models.Follow{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete follow relations"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusAccepted, response)
}

func FollowUserService(c *gin.Context, userID any, follow serializers.Following) {
	var followerUser models.UserDetails
	if err := database.GORM_DB.Where("id = ?", userID).First(&followerUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You deleted your account"})
		return
	}

	var followingUser models.UserDetails
	if err := database.GORM_DB.Where("id = ?", follow.Following).First(&followingUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User you trying to follow does not exist"})
		return
	}

	var existingFollow models.Follow
	if err := database.GORM_DB.Where("follower = ? AND following = ?", userID, follow.Following).First(&existingFollow).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are already following this user"})
		return
	}

	data := models.Follow{
		Follower:  userID.(uint),
		Following: follow.Following,
	}

	if err := database.GORM_DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func GetAllActiveUsersService(c *gin.Context, allusers []serializers.AllUsers) {
	if err := database.GORM_DB.Model(&models.UserDetails{}).
		Select("id, email").
		Find(&allusers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Users": allusers})
}

func GetFollowersService(c *gin.Context, userID any) {
	var follows []models.Follow
	if err := database.GORM_DB.Where("following = ?", userID).Find(&follows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve followers list"})
		return
	}

	var userfollowers []serializers.Followoutput
	for _, follow := range follows {
		var user serializers.Followoutput
		if err := database.GORM_DB.Model(&models.UserDetails{}).Where("id = ?", follow.Follower).First(&user).Error; err == nil {
			userfollowers = append(userfollowers, user)
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"followers": userfollowers})
}

func GetFollowingService(c *gin.Context, userID any) {
	var follows []models.Follow
	if err := database.GORM_DB.Where("follower = ?", userID).Find(&follows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve following list"})
		return
	}

	var followingUsers []serializers.Followoutput
	for _, follow := range follows {
		var user serializers.Followoutput
		if err := database.GORM_DB.Model(models.UserDetails{}).Where("id = ?", follow.Following).First(&user).Error; err == nil {
			followingUsers = append(followingUsers, user)
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"following": followingUsers})
}

func GetuserService(c *gin.Context, userID any) {

	var user models.UserDetails

	if err := database.GORM_DB.Preload("OfficeDetails").Preload("ResidentialDetails").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.BuildUpdateResponse(user)

	c.JSON(http.StatusAccepted, response)
}

func UnfollowUserService(c *gin.Context, unfollow serializers.Following, userID any) {
	var follow models.Follow
	if err := database.GORM_DB.Where("follower = ? AND following = ?", userID, unfollow.Following).First(&follow).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not following this user"})
		return
	}

	if err := database.GORM_DB.Delete(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Successfully unfollowed the user"})
}

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

func UpdateUserService(c *gin.Context, updateInput serializers.UserUpdateInput, userID any) {
	updatedFields := map[string]interface{}{
		"first_name":     updateInput.FirstName,
		"last_name":      updateInput.LastName,
		"gender":         updateInput.Gender,
		"date_of_birth":  updateInput.DateOfBirth,
		"marital_status": updateInput.MaritalStatus,
	}

	if err := database.GORM_DB.Model(&models.UserDetails{}).
		Where("id = ?", userID).
		Updates(updatedFields).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	var user models.UserDetails

	if err := database.GORM_DB.Preload("OfficeDetails").Preload("ResidentialDetails").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": userID})
		return
	}

	response := serializers.BuildUpdateResponse(user)

	c.JSON(http.StatusCreated, response)
}
