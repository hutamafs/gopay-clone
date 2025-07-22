package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type CreateMerchantRequest struct {
	Name              string  `json:"name" validate:"required"`
	Email             string  `json:"email" validate:"required,email"`
	Password          string  `json:"password" validate:"required,min=6"`
	Phone             string  `json:"phone" validate:"required,len=10"`
	Type              string  `json:"user_type"`
	Location          string  `json:"location" validate:"required"`
	MerchantName      string  `json:"merchant_name" validate:"required"`
	Description       string  `json:"description"`
	MerchantPhone     *string `json:"merchant_phone"`
	Category          string  `json:"category"`
	OpenHour          string  `json:"open_hour" validate:"required"`
	ClosedHour        string  `json:"closed_hour" validate:"required"`
	ProfilePictureURL string  `json:"profile_picture_url"`
	MerchantLogoURL   string  `json:"merchant_logo_url"`
}

type UpdateMerchantRequest struct {
	Location        string `json:"location" validate:"required"`
	MerchantName    string `json:"merchant_name" validate:"required"`
	Description     string `json:"description"`
	MerchantPhone   string `json:"merchant_phone"`
	Category        string `json:"category"`
	OpenHour        string `json:"open_hour" validate:"required"`
	ClosedHour      string `json:"closed_hour" validate:"required"`
	MerchantLogoURL string `json:"merchant_logo_url"`
}

func ValidateCreateMerchant(req *CreateMerchantRequest) error {
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
	if req.Type != "merchant" {
		return errors.New("user type must be merchant")
	}
	if err := validateEmptyString(req.MerchantName, "merchant name"); err != nil {
		return err
	}
	if err := validateEmptyString(req.Location, "merchant location"); err != nil {
		return err
	}
	if err := validateOpenClosedHours(req.OpenHour, req.ClosedHour); err != nil {
		return err
	}
	if req.MerchantPhone != nil && *req.MerchantPhone != "" {
		if err := validatePhone(*req.MerchantPhone); err != nil {
			return errors.New("merchant phone number is not valid")
		}
	}

	return nil
}

func ValidateUpdateMerchant(req *UpdateMerchantRequest) error {
	if err := validateEmptyString(req.MerchantName, "merchant name"); err != nil {
		return err
	}
	if err := validateEmptyString(req.Location, "merchant location"); err != nil {
		return err
	}
	if err := validateOpenClosedHours(req.OpenHour, req.ClosedHour); err != nil {
		return err
	}

	return nil
}

func validateEmptyString(str string, key string) error {
	if strings.TrimSpace(str) == "" {
		e := fmt.Sprintf("%v is required", key)
		return errors.New(e)
	}
	return nil
}
func validateOpenClosedHours(openHour, closedHour string) error {
	if strings.TrimSpace(openHour) == "" {
		return errors.New("open hour is required")
	}
	if strings.TrimSpace(closedHour) == "" {
		return errors.New("closed hour is required")
	}

	// Validate time format (HH:MM)
	if err := validateTimeFormat(openHour); err != nil {
		return errors.New("open hour " + err.Error())
	}
	if err := validateTimeFormat(closedHour); err != nil {
		return errors.New("closed hour " + err.Error())
	}

	// Simple time comparison (assuming same day)
	if openHour >= closedHour {
		return errors.New("open hour must be before closed hour")
	}
	return nil
}

func validateTimeFormat(timeStr string) error {
	// Simple validation for HH:MM format
	if !regexp.MustCompile(`^([01]?[0-9]|2[0-3]):[0-5][0-9]$`).MatchString(timeStr) {
		return errors.New("must be in HH:MM format (e.g., 09:00)")
	}
	return nil
}
