package validator

import (
	"errors"
	"gopay-clone/models"
)

var validTransactionTypes = map[models.TransactionType]bool{
	models.Payment:  true,
	models.Transfer: true,
	models.Topup:    true,
	models.Cashback: true,
}

var validTransactionCategory = map[models.TransactionCategory]bool{
	models.Food:          true,
	models.Transport:     true,
	models.Bills:         true,
	models.Entertainment: true,
	models.Other:         true,
}

type CreateTransactionRequest struct {
	Amount            float64                    `json:"amount" gorm:"not null"`
	SenderAccountId   uint                       `json:"sender_id" gorm:"not null"`
	ReceiverAccountId uint                       `json:"receiver_id" gorm:"not null"`
	Type              models.TransactionType     `json:"type" gorm:"not null"`
	Category          models.TransactionCategory `json:"category" gorm:"not null"`
}

func isValidType(t models.TransactionType) bool {
	return validTransactionTypes[t]
}

func isValidCategory(t models.TransactionCategory) bool {
	return validTransactionCategory[t]
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

	if !isValidType(req.Type) {
		return errors.New("not a valid type")
	}

	if !isValidCategory(req.Category) {
		return errors.New("not a valid category")
	}

	return nil
}
