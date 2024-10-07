package serializers

import (
	"app/models"
	"time"
)

type UserUpdateInput struct {
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	DateOfBirth   time.Time      `json:"date_of_birth" gorm:"not null;type:date"`
	Gender        models.Genders `json:"gender"`
	MaritalStatus models.Marital `json:"marital_status"`
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
type ResidentialDetails struct {
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	ContactNo1 string `json:"contact_no_1" binding:"required"`
	ContactNo2 string `json:"contact_no_2"`
}
type User struct {
	Email              string             `json:"email"`
	Password           string             `json:"password"`
	FirstName          string             `json:"first_name"`
	LastName           string             `json:"last_name"`
	Gender             models.Genders     `json:"gender"`
	MaritalStatus      models.Marital     `json:"marital_status"`
	DateOfBirth        time.Time          `json:"date_of_birth"`
	ResidentialDetails ResidentialDetails `json:"residential_details"`
	OfficeDetails      OfficeDetails      `json:"office_details"`
}

type OfficeDetails struct {
	EmployeeCode string `json:"employee_code" binding:"required"`
	Address      string `json:"address" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	ContactNo    string `json:"contact_no" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Name         string `json:"name" binding:"required"`
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

type AllUsers struct {
	ID    uint   `json:"user_id"`
	Email string `json:"email"`
}

type PasswordUpdateInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type Updateresponseuser struct {
	UserID             uint               `json:"user_id"`
	Email              string             `json:"email"`
	LastModified       string             `json:"last_modified"`
	FirstName          string             `json:"first_name"`
	LastName           string             `json:"last_name"`
	DateOfBirth        time.Time          `json:"date_of_birth"`
	Gender             models.Genders     `json:"gender"`
	MaritalStatus      models.Marital     `json:"marital_status"`
	ResidentialDetails ResidentialDetails `json:"residential_details"`
	OfficeDetails      OfficeDetails      `json:"office_details"`
}

func BuildUpdateResponse(userDetails models.UserDetails) Updateresponseuser {
	return Updateresponseuser{
		UserID:             userDetails.ID,
		Email:              userDetails.Email,
		LastModified:       userDetails.UpdatedAt.Format(time.RFC3339),
		FirstName:          userDetails.FirstName,
		LastName:           userDetails.LastName,
		DateOfBirth:        userDetails.DateOfBirth,
		Gender:             userDetails.Gender,
		MaritalStatus:      userDetails.MaritalStatus,
		ResidentialDetails: BuildUpdatedResidentialDetails(userDetails.ResidentialDetails),
		OfficeDetails:      BuildUpdatedOfficeDetails(userDetails.OfficeDetails),
	}
}

func BuildUpdatedResidentialDetails(residentialDetails models.ResidentialDetails) ResidentialDetails {
	return ResidentialDetails{
		Address:    residentialDetails.Address,
		City:       residentialDetails.City,
		State:      residentialDetails.State,
		Country:    residentialDetails.Country,
		ContactNo1: residentialDetails.ContactNo1,
		ContactNo2: residentialDetails.ContactNo2,
	}
}

func BuildUpdatedOfficeDetails(officeDetails models.OfficeDetails) OfficeDetails {
	return OfficeDetails{
		EmployeeCode: officeDetails.EmployeeCode,
		Address:      officeDetails.Address,
		City:         officeDetails.City,
		State:        officeDetails.State,
		Country:      officeDetails.Country,
		ContactNo:    officeDetails.ContactNo,
		Email:        officeDetails.Email,
		Name:         officeDetails.Name,
	}
}

type UserResponse struct {
	UserID             uint               `json:"user_id"`
	Email              string             `json:"email"`
	LastModified       string             `json:"last_modified"`
	FirstName          string             `json:"first_name"`
	LastName           string             `json:"last_name"`
	DateOfBirth        time.Time          `json:"date_of_birth"`
	Gender             models.Genders     `json:"gender"`
	MaritalStatus      models.Marital     `json:"marital_status"`
	ResidentialDetails ResidentialDetails `json:"residential_details"`
	OfficeDetails      OfficeDetails      `json:"office_details"`
	Token              Token              `json:"token"`
}

type Token struct {
	Key        string `json:"key"`
	ExpiryTime string `json:"expiry_time"`
}

func BuildUserResponse(userDetails models.UserDetails, residentialDetails models.ResidentialDetails, officeDetails models.OfficeDetails, token string, expiryTime string) UserResponse {
	return UserResponse{
		UserID:             userDetails.ID,
		Email:              userDetails.Email,
		LastModified:       userDetails.UpdatedAt.Format(time.RFC3339),
		FirstName:          userDetails.FirstName,
		LastName:           userDetails.LastName,
		DateOfBirth:        userDetails.DateOfBirth,
		Gender:             userDetails.Gender,
		MaritalStatus:      userDetails.MaritalStatus,
		ResidentialDetails: BuildResidentialDetails(residentialDetails),
		OfficeDetails:      BuildOfficeDetails(officeDetails),
		Token:              BuildToken(token, expiryTime),
	}
}

func BuildResidentialDetails(residentialDetails models.ResidentialDetails) ResidentialDetails {
	return ResidentialDetails{
		Address:    residentialDetails.Address,
		City:       residentialDetails.City,
		State:      residentialDetails.State,
		Country:    residentialDetails.Country,
		ContactNo1: residentialDetails.ContactNo1,
		ContactNo2: residentialDetails.ContactNo2,
	}
}

func BuildOfficeDetails(officeDetails models.OfficeDetails) OfficeDetails {
	return OfficeDetails{
		EmployeeCode: officeDetails.EmployeeCode,
		Address:      officeDetails.Address,
		City:         officeDetails.City,
		State:        officeDetails.State,
		Country:      officeDetails.Country,
		ContactNo:    officeDetails.ContactNo,
		Email:        officeDetails.Email,
		Name:         officeDetails.Name,
	}
}

func BuildToken(token string, expiryTime string) Token {
	return Token{
		Key:        token,
		ExpiryTime: expiryTime,
	}
}

type Followinput struct {
	Id        uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Followoutput struct {
	Id        uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Following struct {
	Following uint `json:"following"`
}

func GetFollowingDetails(foundUser models.UserDetails) Followoutput {
	return Followoutput{
		Id:        foundUser.ID,
		FirstName: foundUser.FirstName,
		LastName:  foundUser.LastName,
	}

}
