package models

import (
	"gorm.io/gorm"
)

type Follow struct {
	*gorm.Model
	Follower  int64 `json:"follower"`
	Following int64 `json:"following"`
}
