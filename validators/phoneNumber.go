package validators

import (
	"regexp"
)

func ValidatePhoneNumber(phone string) bool {

	regex := `^\+91[6-9]\d{9}$`

	re := regexp.MustCompile(regex)
	return re.MatchString(phone)
}
