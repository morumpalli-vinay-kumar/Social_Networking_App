package models

type Follow struct {
	ID        uint `gorm:"primaryKey"`
	Follower  uint `gorm:"not null"`
	Following uint `gorm:"not null"`

	FollowerUser  UserDetails `gorm:"foreignkey:Follower;references:ID;constraint:OnDelete:CASCADE;"`
	FollowingUser UserDetails `gorm:"foreignkey:Following;references:ID;constraint:OnDelete:CASCADE;"`
}
