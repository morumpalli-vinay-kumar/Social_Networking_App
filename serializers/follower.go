package serializers

import "app/models"

type Followinput struct {
	Id        uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Followoutput struct {
	Id        uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Following struct {
	Following uint `json:"following"`
}

func GetFollowingDetails(foundUser models.UserDetails) Followoutput {
	return Followoutput{
		Id:        foundUser.ID,
		FirstName: foundUser.FirstName,
		LastName:  foundUser.LastName,
	}

}
