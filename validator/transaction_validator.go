package validator

import (
	"errors"
)

type CreateTransactionRequest struct {
	Amount            float64 `json:"amount" gorm:"not null"`
	SenderAccountId   uint    `json:"sender_id" gorm:"not null"`
	ReceiverAccountId uint    `json:"receiver_id" gorm:"not null"`
}

func ValidateCreateTransaction(req *CreateTransactionRequest) error {

	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if req.SenderAccountId == 0 || req.ReceiverAccountId == 0 {
		return errors.New("both sender and receiver ids cannot be empty")
	}

	if req.SenderAccountId == req.ReceiverAccountId {
		return errors.New("both sender and receiver can not be same")
	}
	return nil
}
