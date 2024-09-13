package validators

import (
	"errors"
	"strings"
)

var Genders = map[string]bool{
	"male":       true,
	"female":     true,
	"non-binary": true,
	"other":      true,
}

func ValidateGender(gender string) error {
	gender = strings.ToLower(gender)

	if _, valid := Genders[gender]; !valid {
		return errors.New("enter proper gender")
	}

	return nil
}
