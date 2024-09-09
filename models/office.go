package models

import (
	"gorm.io/gorm"
)

type Office struct {
	*gorm.Model
	UserID      int64  `json:"user_id"`
	OfficeID    int64  `json:"office_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AddressID   int64  `json:"address_id"`
	PhoneNumber string `json:"phone_number"`
}
