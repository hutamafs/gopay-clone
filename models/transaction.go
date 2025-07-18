package models

type Transaction struct {
	BaseModel
	Amount            float64 `json:"amount" gorm:"not null"`
	SenderAccountId   uint    `json:"sender_account_id" gorm:"not null"`
	SenderAccount     Account `json:"sender_account"`
	ReceiverAccountId uint    `json:"receiver_account_id" gorm:"not null"`
	ReceiverAccount   Account `json:"receiver_account"`
	Category          string  `json:"category"`
	Type              string  `json:"type"`
	Status            string  `json:"status" gorm:"default:pending"`
	QrCodeID          *uint   `json:"qr_code_id,omitempty"`
	QrCode            *QrCode `json:"qr_code"`
}
