package utils

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 48

func ValidatePassword(password string) error {
	return passwordvalidator.Validate(password, minEntropyBits)
}
