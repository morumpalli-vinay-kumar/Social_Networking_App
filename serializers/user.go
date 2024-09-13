package serializers

import (
	"app/models"
	"time"
)

type UserUpdateInput struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	MaritalStatus string `json:"marital_status"`
}

type Logininput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Loginresponse struct {
	Id           uint
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LastModified string `json:"last_modified"`
}

type User struct {
	Email              string             `json:"email" validate:"required,email" gorm:"unique;not null"`
	Password           string             `json:"password" validate:"required,min=8"`
	FirstName          string             `json:"first_name" validate:"required"`
	LastName           string             `json:"last_name" validate:"required"`
	DateOfBirth        string             `json:"date_of_birth" validate:"required"`
	Gender             string             `json:"gender" validate:"required"`
	MaritalStatus      string             `json:"marital_status" validate:"required"`
	ResidentialDetails ResidentialDetails `json:"residential_details" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
	OfficeDetails      OfficeDetails      `json:"office_details" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
}

func Loginoutput(foundUser models.UserDetails) Loginresponse {
	return Loginresponse{
		Id:           foundUser.ID,
		Email:        foundUser.Email,
		FirstName:    foundUser.FirstName,
		LastName:     foundUser.LastName,
		LastModified: foundUser.UpdatedAt.Format(time.RFC3339),
	}
}
