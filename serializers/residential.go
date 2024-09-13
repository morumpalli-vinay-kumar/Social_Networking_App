package serializers

type ResidentialDetails struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	ContactNo1 string `json:"contact_no_1"`
	ContactNo2 string `json:"contact_no_2"`
}
