package models

import "time"

type VehicleType string
type UserType string

const (
	Car        VehicleType = "car"
	MotorCycle VehicleType = "motorcycle"
)

const (
	Consumer UserType = "consumer"
	Driver   UserType = "driver"
	Merchant UserType = "merchant"
)

type User struct {
	BaseModel
	Name              string    `json:"name" gorm:"not null"`
	Email             string    `json:"email" gorm:"unique;not null"`
	Phone             int       `json:"phone"`
	Type              UserType  `json:"user_type" gorm:"default:Consumer"`
	Password          string    `json:"-"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	Accounts          []Account `json:"accounts,omitempty"` // we dont have to put gorm fk here because we haev UserId at account, so gorm will assume it is the fk

	// contacts i created
	Contacts []Contact `json:"contacts,omitempty" gorm:"foreignKey:OwnerID"` // the reason we put it here foreignkey ownerId is because this is an user struct and at contact, we have it as owner, not UserId, so gorm needs precise fk explicitly.
}

type DriverProfile struct {
	BaseModel
	UserId            uint        `json:"user_id,omitempty" gorm:"foreignKey:UserId"`
	User              User        `json:"-"`
	LicenseNumber     string      `json:"license_number" gorm:"unique; not null"`
	LicensePictureURL string      `json:"license_picture_url"`
	VehiclePlate      string      `json:"vehicle_plate" gorm:"unique; not null"`
	VehicleType       VehicleType `json:"vehicle_type" gorm:"default:motorcycle;not null"`
	Rating            float64     `json:"rating" gorm:"default:0;not null"`
	CurrentLocation   string      `json:"current_location"`
	Status            string      `json:"status" gorm:"default:offline"` // offline online suspend
	IsVerified        bool        `json:"is_verified" gorm:"default:false"`
}

type MerchantProfile struct {
	BaseModel
	UserId          uint   `json:"user_id,omitempty" gorm:"foreignKey:UserId"`
	User            User   `json:"-"`
	Location        string `json:"location" gorm:"not null"`
	MerchantName    string `json:"merchant_name" gorm:"not null"`
	Description     string `json:"description" gorm:"not null"`
	MerchantPhone   int    `json:"merchant_phone"`
	Category        string `json:"category"`
	OpenHour        time.Time
	ClosedHour      time.Time
	Rating          float64 `json:"rating" gorm:"default:0;not null"`
	MerchantLogoURL string  `json:"merchant_logo_url"`
	Menu            []MenuItem
}

type LoggedinUser struct {
	Email    string `json:"email" `
	Password string `json:"-"`
}
