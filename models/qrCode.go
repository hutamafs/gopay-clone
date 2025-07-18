package models

import "time"

type QrCode struct {
	BaseModel
	AccountID uint      `json:"account_id" gorm:"not null"`
	Account   Account   `json:"account,omitempty"`
	Amount    float64   `json:"amount" gorm:"not null"`
	URL       string    `json:"url" gorm:"not null"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
	ExpiresAt time.Time `json:"expires_at"`
}
