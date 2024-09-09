package models

import (
	"gorm.io/gorm"
)

type Office struct {
	*gorm.Model
	OfficeID      int64  `json:"office_id"`
	Name          string `json:"name"`
	PersonalEmail string `json:"personalmail"`
	OfficeMail    string `json:"officemail"`
	City          string `json:"city"`
	State         string `json:"state"`
	Pincode       string `json:"pincode"`
	Country       string `json:"country"`
	PhoneNumber   string `json:"phone_number"`
}
