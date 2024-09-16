package models

import "gorm.io/gorm"

type OfficeDetails struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	EmployeeCode string `gorm:"not null"`
	Address      string `gorm:"not null"`
	City         string `gorm:"not null"`
	State        string `gorm:"not null"`
	Country      string `gorm:"not null"`
	ContactNo    string `gorm:"not null"`
	Email        string `gorm:"unique:not null"`
	Name         string `gorm:"not null"`
}
