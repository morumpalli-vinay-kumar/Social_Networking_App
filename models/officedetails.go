package models

type OfficeDetails struct {
	UserID       uint   `gorm:"not null"`
	EmployeeCode string `gorm:"unique;not null"`
	Address      string `gorm:"not null"`
	City         string `gorm:"not null"`
	State        string `gorm:"not null"`
	Country      string `gorm:"not null"`
	ContactNo    string `gorm:"not null"`
	Email        string `gorm:"not null"`
	Name         string `gorm:"not null"`
}
