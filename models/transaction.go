package models

type TransactionType string
type TransactionCategory string
type TransactionStatus string
type ServiceType string

const (
	Payment  TransactionType = "payment"
	Transfer TransactionType = "transfer"
	Topup    TransactionType = "topup"
	Cashback TransactionType = "cashback"
)

const (
	Food          TransactionCategory = "food"
	Transport     TransactionCategory = "transport"
	Bills         TransactionCategory = "bills"
	Entertainment TransactionCategory = "entertainment"
	TransferCat   TransactionCategory = "transfer" // for example moving balance from 1 to another user
	Other         TransactionCategory = "other"
)

const (
	TransactionPending   TransactionStatus = "pending"
	TransactionCompleted TransactionStatus = "completed"
	TransactionFailed    TransactionStatus = "failed"
	TransactionCancelled TransactionStatus = "cancelled"
)

const (
	ServiceFood ServiceType = "food"
	ServiceRide ServiceType = "ride"
	ServiceNone ServiceType = "none" // for regular transfers
)

type Transaction struct {
	BaseModel
	Amount            float64             `json:"amount" gorm:"not null"`
	SenderAccountId   uint                `json:"sender_account_id" gorm:"not null;index:idx_sender_account"`
	SenderAccount     Account             `json:"sender_account" gorm:"foreignKey:SenderAccountId"`
	ReceiverAccountId uint                `json:"receiver_account_id" gorm:"not null;index:idx_receiver_account"`
	ReceiverAccount   Account             `json:"receiver_account" gorm:"foreignKey:ReceiverAccountId"`
	Category          TransactionCategory `json:"category" gorm:"default:other;index:idx_category"`
	Type              TransactionType     `json:"type" gorm:"default:payment;index:idx_type"`
	Status            TransactionStatus   `json:"status" gorm:"default:pending;index:idx_status"`
	QrCodeID          *uint               `json:"qr_code_id,omitempty"`
	QrCode            *QrCode             `json:"qr_code,omitempty" gorm:"foreignKey:QrCodeID"`
	Description       string              `json:"description,omitempty"`
	ServiceType       ServiceType         `json:"service_type" gorm:"default:none;index:idx_service_type"`
	ServiceID         *uint               `json:"service_id,omitempty" gorm:"index:idx_service_id"` // optional because it might be just a transfer
}
