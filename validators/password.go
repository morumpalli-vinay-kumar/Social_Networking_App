package validators

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func ValidatePassword(pass string) error {
	const minEntropyBits = 62
	return passwordvalidator.Validate(pass, minEntropyBits)
}
