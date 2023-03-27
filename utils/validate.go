package utils

import (
	"errors"
	"net/mail"
	"projectsync/dto"
)

func Validate(T any) error {

	switch v := T.(type) {

	case dto.Register:
		return validateRegisterDTO(v)
	}
	return nil
}

func validateRegisterDTO(register dto.Register) error {
	if register.ConfirmPassword != register.ConfirmPassword {
		return errors.New("password don't match")
	}
	if register.Password == "" {
		return errors.New("password is empty")
	}
	if _, err := mail.ParseAddress(register.Email); err != nil {
		return errors.New("invalid email address")
	}
	return nil
}
