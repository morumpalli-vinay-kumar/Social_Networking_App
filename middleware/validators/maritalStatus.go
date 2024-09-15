package validators

import (
	"app/models"
	"errors"
	"strings"
)

func ValidateMaritalStatus(m models.Marital) error {
	switch strings.ToLower(string(m)) {
	case string(models.Married), string(models.Single), string(models.Divorced), string(models.Widowed):
		return nil
	}
	return errors.New("invalid gender")
}
