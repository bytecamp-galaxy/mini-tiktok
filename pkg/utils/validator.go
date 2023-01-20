package utils

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 60

func ValidatePassword(password string) error {
	return passwordvalidator.Validate(password, minEntropyBits)
}
