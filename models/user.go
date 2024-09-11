package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        int64  `json:"user_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PersonalEmail string `json:"personal_mail"`
	Password      string `json:"password"`
	Gender        string `json:"gender"`
	PhoneNumber   string `json:"phone_number"`
	DOB           string `json:"dob"`
	City          string `json:"city"`
	State         string `json:"state"`
	Pincode       string `json:"pincode"`
	Country       string `json:"country"`
	IsActive      bool   `json:"is_active"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}
