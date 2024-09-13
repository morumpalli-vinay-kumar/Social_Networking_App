package serializers

import (
	"app/models"
	"time"
)

type Updateresponseuser struct {
	UserID             uint               `json:"user_id"`
	Email              string             `json:"email"`
	LastModified       string             `json:"last_modified"`
	FirstName          string             `json:"first_name"`
	LastName           string             `json:"last_name"`
	DateOfBirth        string             `json:"date_of_birth"`
	Gender             string             `json:"gender"`
	MaritalStatus      string             `json:"marital_status"`
	ResidentialDetails ResidentialDetails `json:"residential_details"`
	OfficeDetails      OfficeDetails      `json:"office_details"`
}

func BuildUpdateResponse(userDetails models.UserDetails, residentialDetails models.ResidentialDetails, officeDetails models.OfficeDetails) Updateresponseuser {
	return Updateresponseuser{
		UserID:             userDetails.ID,
		Email:              userDetails.Email,
		LastModified:       userDetails.UpdatedAt.Format(time.RFC3339),
		FirstName:          userDetails.FirstName,
		LastName:           userDetails.LastName,
		DateOfBirth:        userDetails.DateOfBirth,
		Gender:             userDetails.Gender,
		MaritalStatus:      userDetails.MaritalStatus,
		ResidentialDetails: BuildUpdatedResidentialDetails(residentialDetails),
		OfficeDetails:      BuildUpdatedOfficeDetails(officeDetails),
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
