package validators

import (
	"app/middleware/serializers"
	"app/models"
	"errors"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func ValidatePhoneNumber(phone string) bool {

	regex := `^\+91[6-9]\d{9}$`

	re := regexp.MustCompile(regex)
	return re.MatchString(phone)
}

func ValidatePassword(pass string) error {
	const minEntropyBits = 50
	return passwordvalidator.Validate(pass, minEntropyBits)
}

func ValidateMaritalStatus(m models.Marital) error {
	switch strings.ToLower(string(m)) {
	case string(models.Married), string(models.Single), string(models.Divorced), string(models.Widowed):
		return nil
	}
	return errors.New("invalid gender")
}

func ValidateGender(g models.Genders) error {
	switch strings.ToLower(string(g)) {
	case string(models.Male), string(models.Female), string(models.Other):
		return nil
	}
	return errors.New("invalid gender")
}

func ValidateEmail(email string) bool {
	return govalidator.IsEmail(email)
}
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
