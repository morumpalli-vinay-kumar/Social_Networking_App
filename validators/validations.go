package validators

import (
	"app/serializers"
	"errors"
)

func ValidationCheck(req serializers.User) error {

	if err := ValidatePassword(req.Password); err != nil {
		return errors.New("please use strong password")
	}

	if check := ValidatePhoneNumber(req.ResidentialDetails.ContactNo1); !check {
		return errors.New("invalid residential phone number")
	}
	if check := ValidatePhoneNumber(req.OfficeDetails.ContactNo); !check {
		return errors.New("invalid office phone number")
	}

	if err := ValidateGender(req.Gender); err != nil {
		return err
	}

	if err := ValidateMaritalStatus(req.MaritalStatus); err != nil {
		return err
	}

	if check := ValidateEmail(req.Email); !check {
		return errors.New("invalid-mail")
	}

	return nil
}
