package serializers

type ResidentialDetails struct {
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	ContactNo1 string `json:"contact_no_1" binding:"required"`
	ContactNo2 string `json:"contact_no_2"`
}
