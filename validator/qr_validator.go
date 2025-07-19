package validator

import (
	"errors"
	"strings"
	"time"
)

type CreateQRRequest struct {
	ReceiverAccountID uint      `json:"receiver_account_id" gorm:"not null"`
	Amount            float64   `json:"amount" gorm:"not null"`
	URL               string    `json:"url" gorm:"not null"`
	ExpiresAt         time.Time `json:"expires_at" gorm:"default:CURRENT_TIMESTAMP + INTERVAL 1 MINUTE"`
}

type ScanQRRequest struct {
	SenderAccountID uint `json:"sender_account_id" gorm:"not null"`
}

func ValidateCreateQR(req *CreateQRRequest) error {

	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if req.ReceiverAccountID == 0 {
		return errors.New("qr receiver id can't be empty")
	}

	if strings.TrimSpace(req.URL) == "" {
		return errors.New("qr url is required")
	}
	return nil
}

func ValidateScanQR(req *ScanQRRequest) error {

	if req.SenderAccountID == 0 {
		return errors.New("sender id can't be empty")
	}

	return nil
}
