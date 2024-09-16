package validators

import (
	"app/serializers"
	"errors"
)

func ValidationChecklogin(loginInput serializers.Logininput) error {
	if check := ValidateEmail(loginInput.Email); !check {
		return errors.New("invalid-mail")
	}
	return nil
}

func ValidationCheckSignup(req serializers.User) error {

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

	if len(req.FirstName) < 3 {
		return errors.New("first name should be minimum of 3 letters")
	}

	if err := ValidateMaritalStatus(req.MaritalStatus); err != nil {
		return err
	}

	if check := ValidateEmail(req.Email); !check {
		return errors.New("invalid-user-mail")
	}

	if check := ValidateEmail(req.OfficeDetails.Email); !check {
		return errors.New("invalid-office-mail")
	}
	return nil
}

func ValidationCheckUpdate(req serializers.UserUpdateInput) error {

	if err := ValidateGender(req.Gender); err != nil {
		return err
	}

	if err := ValidateMaritalStatus(req.MaritalStatus); err != nil {
		return err
	}
	return nil

}
