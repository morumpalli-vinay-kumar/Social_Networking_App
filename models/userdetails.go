package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDetails struct {
	gorm.Model
	Email              string             `gorm:"unique;not null"`
	Password           string             `gorm:"not null"`
	FirstName          string             `gorm:"not null"`
	LastName           string             `gorm:"not null"`
	DateOfBirth        string             `gorm:"not null"`
	Gender             string             `gorm:"not null"`
	MaritalStatus      string             `gorm:"not null"`
	ResidentialDetails ResidentialDetails `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
	OfficeDetails      OfficeDetails      `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
}

func (u *UserDetails) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *UserDetails) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}
