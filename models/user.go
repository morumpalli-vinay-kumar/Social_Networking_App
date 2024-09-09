package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	UserID      int64     `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AddressID   int64     `json:"address_id"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	DOB         string    `json:"dob"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
