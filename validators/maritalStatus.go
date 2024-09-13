package validators

import (
	"errors"
	"strings"
)

var MaritalStatuses = map[string]bool{
	"single":   true,
	"married":  true,
	"divorced": true,
	"widowed":  true,
}

func ValidateMaritalStatus(status string) error {
	status = strings.ToLower(status)

	if valid := MaritalStatuses[status]; !valid {
		return errors.New("enter proper marital status")
	}

	return nil
}
