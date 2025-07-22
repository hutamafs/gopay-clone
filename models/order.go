package models

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderConfirmed OrderStatus = "confirmed"
	OrderCooking   OrderStatus = "cooking"
	OrderDelivery  OrderStatus = "delivery"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

type Order struct {
	BaseModel
	UserID          uint            `json:"user_id" gorm:"not null;index:idx_user_id"`         // who orders
	MerchantID      uint            `json:"merchant_id" gorm:"not null;index:idx_merchant_id"` // where did the user order
	User            User            `json:"-" gorm:"foreignKey:UserID"`
	Merchant        MerchantProfile `json:"-" gorm:"foreignKey:MerchantID"`
	DriverID        *uint           `json:"driver_id,omitempty" gorm:"index:idx_driver_id"` //pointer because driver assigned later
	Driver          *DriverProfile  `json:"driver,omitempty" gorm:"foreignKey:DriverID"`
	Items           []OrderItem     `json:"items" gorm:"foreignKey:OrderID"` // Added FK
	TotalAmount     float64         `json:"total_amount"`
	DeliveryFee     float64         `json:"delivery_fee"`
	Status          OrderStatus     `json:"status" gorm:"default:pending;index:idx_status"` // Changed to enum
	DeliveryAddress string          `json:"delivery_address" gorm:"not null"`
	TransactionID   *uint           `json:"transaction_id,omitempty"` // pointer because transaction will be created when payment is processed (usually when order moves from pending to confirmed)
	Transaction     *Transaction    `json:"transaction,omitempty" gorm:"foreignKey:TransactionID"`
}
