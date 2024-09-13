package models

type ResidentialDetails struct {
	UserID     uint   `gorm:"not null"`
	Address    string `gorm:"not null"`
	City       string `gorm:"not null"`
	State      string `gorm:"not null"`
	Country    string `gorm:"not null"`
	ContactNo1 string `gorm:"not null"`
	ContactNo2 string
}
