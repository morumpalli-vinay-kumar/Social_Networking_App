package models

import (
	"gorm.io/gorm"
)

type Address struct {
	*gorm.Model
	AddressID int64  `json:"address_id"`
	City      string `json:"city"`
	State     string `json:"state"`
	Pincode   string `json:"pincode"`
	Country   string `json:"country"`
}
