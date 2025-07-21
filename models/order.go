package models

type Order struct {
	BaseModel
	UserID          uint            `json:"user_id" gorm:"not null"`     // who orders
	MerchantID      uint            `json:"merchant_id" gorm:"not null"` // where did the user order
	User            User            `json:"user" gorm:"foreignKey:UserID"`
	Merchant        MerchantProfile `json:"merchant" gorm:"foreignKey:MerchantID"`
	DriverID        uint            `json:"driver_id" gorm:"not null"` // who sends
	Driver          DriverProfile   `json:"driver" gorm:"foreignKey:DriverID"`
	Items           []OrderItem     `json:"items"`
	TotalAmount     float64         `json:"total_amount"`
	DeliveryFee     float64         `json:"delivery_fee"`
	Status          string          `json:"status" gorm:"default:pending"` // "pending", "confirmed", "cooking", "delivery", "completed", "cancelled"
	DeliveryAddress string          `json:"delivery_address" gorm:"not null"`
	TransactionID   uint            `json:"transaction_id"`
	Transaction     Transaction     `json:"transaction"`
}
