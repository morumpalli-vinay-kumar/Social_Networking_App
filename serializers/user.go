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
	Email              string             `json:"email"`
	Password           string             `json:"password"`
	FirstName          string             `json:"first_name"`
	LastName           string             `json:"last_name"`
	DateOfBirth        string             `json:"date_of_birth"`
	Gender             string             `json:"gender"`
	MaritalStatus      string             `json:"marital_status"`
	ResidentialDetails ResidentialDetails `json:"residential_details"`
	OfficeDetails      OfficeDetails      `json:"office_details"`
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
