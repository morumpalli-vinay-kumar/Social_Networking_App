package serializers

type OfficeDetails struct {
	EmployeeCode string `json:"employee_code" binding:"required"`
	Address      string `json:"address" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	ContactNo    string `json:"contact_no" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Name         string `json:"name" binding:"required"`
}
