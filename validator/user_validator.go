package validator

import (
	"errors"
	"regexp"
	"strings"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"omitempty"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

func ValidateCreateUser(req *CreateUserRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}

func ValidateUpdateUser(req *UpdateUserRequest) error {
	if req.Name != "" && strings.TrimSpace(req.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if req.Password != "" && len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
