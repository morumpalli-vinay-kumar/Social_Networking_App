package models

type Follow struct {
	ID        uint `gorm:"primaryKey"`
	Follower  uint `gorm:"not null"`
	Following uint `gorm:"not null"`
}
