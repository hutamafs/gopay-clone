package validator

import (
	"errors"
	"gopay-clone/models"
	"strings"
)

var validVehicleTypes = map[models.VehicleType]bool{
	models.MotorCycle: true,
	models.Car:        true,
}

var validDriverStatuses = map[string]bool{
	"online":  true,
	"sending": true,
	"offline": true,
	"suspend": true,
}

type CreateDriverRequest struct {
	Name              string             `json:"name" validate:"required"`
	Email             string             `json:"email" validate:"required,email"`
	Password          string             `json:"password" validate:"required,min=6"`
	Phone             string             `json:"phone" validate:"required,len=10"`
	Type              string             `json:"user_type"`
	LicenseNumber     string             `json:"license_number" validate:"required"`
	LicensePictureURL string             `json:"license_picture_url"`
	VehiclePlate      string             `json:"vehicle_plate" validate:"required"`
	VehicleType       models.VehicleType `json:"vehicle_type" validate:"required"`
	ProfilePictureURL string             `json:"profile_picture_url"`
	CurrentLocation   string             `json:"current_location"`
}

type UpdateDriverRequest struct {
	LicenseNumber     string             `json:"license_number"`
	LicensePictureURL string             `json:"license_picture_url"`
	VehiclePlate      string             `json:"vehicle_plate"`
	VehicleType       models.VehicleType `json:"vehicle_type"`
	CurrentLocation   string             `json:"current_location"`
}

type UpdateDriverStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

type UpdateDriverLocationRequest struct {
	CurrentLocation string `json:"current_location" validate:"required"`
}

func ValidateCreateDriver(req *CreateDriverRequest) error {
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
	if req.Type != "driver" {
		return errors.New("user type must be driver")
	}
	if strings.TrimSpace(req.LicenseNumber) == "" {
		return errors.New("license number cannot be empty")
	}
	if strings.TrimSpace(req.VehiclePlate) == "" {
		return errors.New("vehicle plate cannot be empty")
	}
	if !isValidVehicleType(req.VehicleType) {
		return errors.New("invalid vehicle type")
	}
	return nil
}

func ValidateUpdateDriver(req *UpdateDriverRequest) error {
	if req.VehicleType != "" && !isValidVehicleType(req.VehicleType) {
		return errors.New("invalid vehicle type")
	}
	return nil
}

func ValidateUpdateDriverStatus(req *UpdateDriverStatusRequest) error {
	if !isValidDriverStatus(req.Status) {
		return errors.New("invalid driver status. Must be: online, offline, or suspend")
	}
	return nil
}

func ValidateUpdateDriverLocation(req *UpdateDriverLocationRequest) error {
	if strings.TrimSpace(req.CurrentLocation) == "" {
		return errors.New("current location cannot be empty")
	}
	return nil
}

func isValidVehicleType(vt models.VehicleType) bool {
	return validVehicleTypes[vt]
}

func isValidDriverStatus(status string) bool {
	return validDriverStatuses[status]
}
