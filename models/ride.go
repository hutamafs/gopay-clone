package models

type Ride struct {
	BaseModel
	UserID          uint          `json:"user_id" gorm:"not null"`
	DriverID        uint          `json:"driver_id"`
	User            User          `json:"user" gorm:"foreignKey:UserID"`
	Driver          DriverProfile `json:"driver" gorm:"foreignKey:DriverID"`
	PickupLocation  string        `json:"pickup_location" gorm:"not null"`
	DropoffLocation string        `json:"dropoff_location" gorm:"not null"`
	VehicleType     VehicleType   `json:"vehicle_type"`
	Status          string        `json:"status" gorm:"default:requested"` // "requested", "accepted", "pickup", "ongoing", "completed", "cancelled"
	Fare            float64       `json:"fare"`
	Distance        float64       `json:"distance"` // in KM
	TransactionID   uint          `json:"transaction_id"`
	Transaction     Transaction   `json:"transaction"`
}
