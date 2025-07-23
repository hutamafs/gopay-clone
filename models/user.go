package models

type VehicleType string
type UserType string
type DriverStatus string

const (
	Car        VehicleType = "car"
	MotorCycle VehicleType = "motorcycle"
)

const (
	Consumer UserType = "consumer"
	Driver   UserType = "driver"
	Merchant UserType = "merchant"
)

const (
	Offline   DriverStatus = "offline"
	Online    DriverStatus = "online"
	Suspended DriverStatus = "suspended"
	Sending   DriverStatus = "sending"
)

type User struct {
	BaseModel
	Name              string    `json:"name" gorm:"not null"`
	Email             string    `json:"email" gorm:"unique;not null;index:idx_email"`
	Phone             string    `json:"phone" gorm:"index:idx_phone"`                          // Changed from int to string
	Type              UserType  `json:"user_type" gorm:"default:consumer;index:idx_user_type"` // Fixed case
	Password          string    `json:"-"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	Accounts          []Account `json:"accounts,omitempty"` // we dont have to put gorm fk here because we haev UserId at account, so gorm will assume it is the fk

	// contacts i created
	Contacts []Contact `json:"contacts,omitempty" gorm:"foreignKey:OwnerID"` // the reason we put it here foreignkey ownerId is because this is an user struct and at contact, we have it as owner, not UserId, so gorm needs precise fk explicitly.
}

type DriverProfile struct {
	BaseModel
	UserId            uint         `json:"user_id,omitempty" gorm:"foreignKey:UserId;index:idx_user_id"`
	User              User         `json:"-"`
	LicenseNumber     string       `json:"license_number" gorm:"unique; not null"`
	LicensePictureURL string       `json:"license_picture_url"`
	VehiclePlate      string       `json:"vehicle_plate" gorm:"unique; not null"`
	VehicleType       VehicleType  `json:"vehicle_type" gorm:"default:motorcycle;not null;index:idx_vehicle_type"`
	Rating            float64      `json:"rating" gorm:"default:0;not null;"`
	CurrentLocation   string       `json:"current_location" gorm:"index:idx_location"`
	Status            DriverStatus `json:"status" gorm:"default:offline;index:idx_status"` // offline online suspend
	IsVerified        bool         `json:"is_verified" gorm:"default:false"`
}

type MerchantProfile struct {
	BaseModel
	UserId          uint       `json:"user_id,omitempty" gorm:"foreignKey:UserId"`
	User            User       `json:"-"`
	Location        string     `json:"location" gorm:"not null;index:idx_merchant_location"`
	MerchantName    string     `json:"merchant_name" gorm:"not null;index:idx_merchant_name"`
	Description     string     `json:"description" gorm:"not null"`
	MerchantPhone   string     `json:"merchant_phone"`
	Category        string     `json:"category" gorm:"index:idx_merchant_category"`
	OpenHour        string     `json:"open_hour"`   // "09:00" format
	ClosedHour      string     `json:"closed_hour"` // "22:00" format
	Rating          float64    `json:"rating" gorm:"default:0;not null"`
	MerchantLogoURL string     `json:"merchant_logo_url"`
	Menu            []MenuItem `json:"menu,omitempty" gorm:"foreignKey:MerchantId"`
}

type LoggedinUser struct {
	Email    string `json:"email" `
	Password string `json:"-"`
}
