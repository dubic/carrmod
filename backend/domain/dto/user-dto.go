package dto

import (
	"errors"
	"net/mail"
)

// user creation request
type UserCreationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreationRequest) Create() []error {
	var errs []error
	if u.Name == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if u.Email == "" {
		errs = append(errs, errors.New("email is required"))
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		errs = append(errs, errors.New("email is not valid"))
	}
	if u.Password == "" {
		errs = append(errs, errors.New("Password is required"))
	}
	if len(u.Password) < 8 || len(u.Password) > 50 {
		errs = append(errs, errors.New("Password must be between 8 to 50 characters"))
	}
	return errs
}
