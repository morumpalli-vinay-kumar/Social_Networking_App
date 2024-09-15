package validators

import (
	"app/models"
	"errors"
	"strings"
)

func ValidateGender(g models.Genders) error {
	switch strings.ToLower(string(g)) {
	case string(models.Male), string(models.Female), string(models.Other):
		return nil
	}
	return errors.New("invalid gender")
}
