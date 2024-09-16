package models

import "gorm.io/gorm"

type ResidentialDetails struct {
	gorm.Model
	UserID     uint   `gorm:"not null"`
	Address    string `gorm:"not null"`
	City       string `gorm:"not null"`
	State      string `gorm:"not null"`
	Country    string `gorm:"not null"`
	ContactNo1 string `gorm:"not null"`
	ContactNo2 string
}
