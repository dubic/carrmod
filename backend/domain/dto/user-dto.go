package dto

import (
	"fmt"
	"net/mail"
	"strings"
	"time"
)

// user creation request
type UserCreationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// user creation response
type UserCreationResponse struct {
	VerificationMailSent bool   `json:"verificationMailSent"`
	Created              bool   `json:"created"`
	Email                string `json:"email"`
	Msg                  string `json:"msg"`
}

// login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// login response
type LoginResponse struct {
	Token         string    `json:"token"`
	Account       Account   `json:"account"`
	LoginTime     time.Time `json:"loggedInAt"`
	Msg           string    `json:"msg"`
	Verified      bool      `json:"verified"`
	Authenticated bool      `json:"authenticated"`
}

// Account
type Account struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *UserCreationRequest) Create() []string {
	var errs []string
	if u.Name == "" {
		errs = append(errs, "name is required")
	}
	if u.Email == "" {
		errs = append(errs, "email is required")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		errs = append(errs, "email is not valid")
	}
	if u.Password == "" {
		errs = append(errs, "password is required")
	}
	if len(u.Password) < 8 || len(u.Password) > 50 {
		errs = append(errs, "password must be between 8 to 50 characters")
	}
	u.Email = strings.ToLower(u.Email)
	return errs
}

func (u UserCreationRequest) Print() string {
	return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

func (l *LoginRequest) Validate() []string {
	var errs []string
	if l.Email == "" {
		errs = append(errs, "email is required")
	}

	if l.Password == "" {
		errs = append(errs, "password is required")
	}
	l.Email = strings.ToLower(l.Email)
	return errs
}
