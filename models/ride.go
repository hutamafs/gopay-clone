package models

type RideStatus string

const (
	RideRequested RideStatus = "requested"
	RideAccepted  RideStatus = "accepted"
	RidePickup    RideStatus = "pickup"
	RideOngoing   RideStatus = "ongoing"
	RideCompleted RideStatus = "completed"
	RideCancelled RideStatus = "cancelled"
)

type Ride struct {
	BaseModel
	UserID          uint           `json:"user_id" gorm:"not null;index:idx_user_id"`
	DriverID        *uint          `json:"driver_id,omitempty" gorm:"index:idx_driver_id"` // optional because driver will be assigned later
	User            User           `json:"user" gorm:"foreignKey:UserID"`
	Driver          *DriverProfile `json:"driver,omitempty" gorm:"foreignKey:DriverID"`
	PickupLocation  string         `json:"pickup_location" gorm:"not null"`
	DropoffLocation string         `json:"dropoff_location" gorm:"not null"`
	VehicleType     VehicleType    `json:"vehicle_type" gorm:"index:idx_vehicle_type"`
	Status          RideStatus     `json:"status" gorm:"default:requested;index:idx_status"`
	Fare            float64        `json:"fare"`
	Distance        float64        `json:"distance"`                 // in KM
	TransactionID   *uint          `json:"transaction_id,omitempty"` // optional because created once the order is completed
	Transaction     *Transaction   `json:"transaction,omitempty" gorm:"foreignKey:TransactionID"`
}
