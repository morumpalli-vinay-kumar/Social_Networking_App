package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Genders string

const (
	Male   Genders = "male"
	Female Genders = "female"
	Other  Genders = "other"
)

type Marital string

const (
	Married  Marital = "married"
	Single   Marital = "single"
	Divorced Marital = "dovorced"
	Widowed  Marital = "widowed"
)

type UserDetails struct {
	gorm.Model
	Email              string `gorm:"unique;not null" validate:"email"`
	Password           string `gorm:"not null" validate:"min=8"`
	FirstName          string `gorm:"not null" validate:"min=3"`
	LastName           string
	DateOfBirth        time.Time          `gorm:"not null;type:date"`
	Gender             Genders            `gorm:"not null"`
	MaritalStatus      Marital            `gorm:"not null"`
	ResidentialDetails ResidentialDetails `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
	OfficeDetails      OfficeDetails      `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
	FollowerUser       Follow             `gorm:"foreignkey:Follower;constraint:OnDelete:CASCADE;"`
	FollowingUser      Follow             `gorm:"foreignkey:Following;constraint:OnDelete:CASCADE;"`
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
