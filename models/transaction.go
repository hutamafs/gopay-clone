package models

type TransactionType string

type TransactionCategory string

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
	Other         TransactionCategory = "other"
)

type Transaction struct {
	BaseModel
	Amount            float64             `json:"amount" gorm:"not null"`
	SenderAccountId   uint                `json:"sender_account_id" gorm:"not null"`
	SenderAccount     Account             `json:"sender_account"`
	ReceiverAccountId uint                `json:"receiver_account_id" gorm:"not null"`
	ReceiverAccount   Account             `json:"receiver_account"`
	Category          TransactionCategory `json:"category" gorm:"default:other"` // food, transport, bills
	Type              TransactionType     `json:"type" gorm:"default:payment"`   // payment, transfer, topup
	Status            string              `json:"status" gorm:"default:pending"`
	QrCodeID          *uint               `json:"qr_code_id,omitempty"`
	QrCode            *QrCode             `json:"qr_code"`
}
