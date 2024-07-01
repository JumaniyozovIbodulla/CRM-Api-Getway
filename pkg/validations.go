package pkg

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateFullName(name string) error {
	if len(strings.Split(name, " ")) < 2 {
		return errors.New("full name's length must be 2")
	}

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be blank")
	}
	if len(password) < 8 || len(password) > 30 {
		return errors.New("password length should be 8 to 30 characters")
	}

	_, err := regexp.MatchString(`^[A-Za-z0-9$_@.#]+$`, password)

	if err != nil {
		return errors.New("password should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}

	return nil
}

func ValidatePhone(phone string) error {
	_, err := regexp.MatchString(`^\+998\d{9}$`, phone)

	if err != nil {
		return errors.New("phone number must be +998")
	}
	return nil
}

func ValidateDay(day string) bool {
	return !strings.EqualFold(day, "sunday") 
}
