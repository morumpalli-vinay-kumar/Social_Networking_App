package validators

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func ValidatePassword(pass string) error {
	const minEntropyBits = 30
	return passwordvalidator.Validate(pass, minEntropyBits)
}
