package models

import "time"

type QrCode struct {
	BaseModel
	ReceiverAccountID uint      `json:"receiver_account_id" gorm:"not null"`
	ReceiverAccount   Account   `json:"receiver_account"`
	Amount            float64   `json:"amount" gorm:"not null"`
	URL               string    `json:"url" gorm:"not null"`
	IsUsed            bool      `json:"is_used" gorm:"default:false"`
	ExpiresAt         time.Time `json:"expires_at"`
}
