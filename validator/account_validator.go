package validator

import (
	"errors"
	"strings"
)

type CreateAccountRequest struct {
	Name    string  `json:"name" gorm:"not null"`
	Balance float64 `json:"balance" gorm:"not null"`
	UserId  uint    `json:"user_id" gorm:"not null"`
}

type UpdateAccountRequest struct {
	Name    string  `json:"name" gorm:"not null"`
	Balance float64 `json:"balance" gorm:"not null"`
}

func ValidateCreateAccount(req *CreateAccountRequest) error {
	if req.Name != "" && strings.TrimSpace(req.Name) == "" {
		return errors.New("account name cannot be empty")
	}

	if req.Balance < 0 {
		return errors.New("balance must be greater than 0")
	}

	if req.UserId == 0 {
		return errors.New("account name cannot be empty")
	}
	return nil
}

func ValidateUpdateAccount(req *UpdateAccountRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("account name cannot be empty")
	}
	if req.Balance < 0 {
		return errors.New("balance must be greater than 0")
	}

	return nil
}
