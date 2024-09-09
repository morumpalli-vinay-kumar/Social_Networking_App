package models

import (
	"gorm.io/gorm"
)

type Id int64

type Follow struct {
	gorm.Model
	UserID    int64 `json:"user_id"`
	Followers []Id  `json:"followers"`
	Following []Id  `json:"following"`
}
