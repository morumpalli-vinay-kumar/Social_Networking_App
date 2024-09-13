package validators

import (
	"app/serializers"
)

func ValidationCheck(req serializers.User) error {

	// if err := ValidatePassword(req.Password); err != nil {
	// 	return errors.New("please use combination on small, capital letters and number in password")
	// }

	// if check := ValidateEmail(req.Email); !check {
	// 	return errors.New("invalid mail")
	// }

	return nil
}
