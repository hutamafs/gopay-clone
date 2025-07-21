package validator

import (
	"errors"
	"regexp"
	"strings"
)

type CreateUserRequest struct {
	Name              string `json:"name" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required,min=6"`
	Phone             string `json:"phone" validate:"required,len=10"`
	Type              string `json:"user_type"`
	ProfilePictureURL string `json:"profile_picture_url"`
}

type UpdateUserRequest struct {
	Name              *string `json:"name" validate:"omitempty"`
	Password          *string `json:"password" validate:"omitempty,min=6"`
	Phone             *string `json:"phone" validate:"omitempty,len=10"`
	ProfilePictureURL *string `json:"profile_picture_url"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ValidateLogin(req *LoginRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return err
	}
	if err := validatePassword(req.Password, true); err != nil {
		return err
	}
	return nil
}

func ValidateCreateUser(req *CreateUserRequest) error {
	if err := validateName(req.Name); err != nil {
		return err
	}
	if err := validateEmail(req.Email); err != nil {
		return err
	}
	if err := validatePassword(req.Password, true); err != nil {
		return err
	}
	if err := validatePhone(req.Phone); err != nil {
		return err
	}
	return nil
}

func ValidateUpdateUser(req *UpdateUserRequest) error {
	if req.Name != nil && *req.Name != "" {
		if err := validateName(*req.Name); err != nil {
			return err
		}
	}
	if req.Password != nil && *req.Password != "" {
		if err := validatePassword(*req.Password, false); err != nil {
			return err
		}
	}
	if req.Phone != nil && *req.Phone != "" {
		if err := validatePhone(*req.Phone); err != nil {
			return err
		}
	}
	return nil
}

func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func validatePassword(pw string, required bool) error {
	if required && strings.TrimSpace(pw) == "" {
		return errors.New("password is required")
	}
	if len(pw) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func validatePhone(phone string) error {
	trimmed := strings.TrimSpace(phone)
	if len(trimmed) != 10 {
		return errors.New("phone must be exactly 10 digits")
	}
	if !regexp.MustCompile(`^\d{10}$`).MatchString(trimmed) {
		return errors.New("phone must be numeric and 10 digits")
	}
	return nil
}
