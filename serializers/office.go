package serializers

type OfficeDetails struct {
	EmployeeCode string `json:"employee_code"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ContactNo    string `json:"contact_no"`
	Email        string `json:"email"`
	Name         string `json:"name"`
}
