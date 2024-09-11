package models

import (
	"gorm.io/gorm"
)

type Office struct {
	gorm.Model
	OfficeID      int64  `json:"office_id"`
	UserID        int64  `json:"user_id"`
	OfficeMail    string `json:"office_mail"`
	OfficeCity    string `json:"office_city"`
	OfficeState   string `json:"office_state"`
	OfficePincode string `json:"office_pincode"`
	OfficeCountry string `json:"office_country"`
}
