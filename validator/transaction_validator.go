package validator

import (
	"errors"
	"gopay-clone/models"
)

var validTransactionStatuses = map[models.TransactionStatus]bool{
	models.TransactionPending:   true,
	models.TransactionCompleted: true,
	models.TransactionFailed:    true,
	models.TransactionCancelled: true,
}

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
	models.TransferCat:   true,
	models.Other:         true,
}

type CreateTransactionRequest struct {
	Amount            float64                     `json:"amount" gorm:"not null"`
	SenderAccountID   uint                        `json:"sender_id" gorm:"not null"`
	ReceiverAccountID uint                        `json:"receiver_id" gorm:"not null"`
	Type              *models.TransactionType     `json:"type,omitempty"`
	Category          *models.TransactionCategory `json:"category,omitempty"`
	Status            *models.TransactionStatus   `json:"status,omitempty"`
	QrCodeID          *uint                       `json:"qr_code_id,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	ServiceID         *uint                       `json:"service_id,omitempty"`
	ServiceType       *models.ServiceType         `json:"service_type,omitempty"`
}

type UpdateTransactionRequest struct {
	Status      *models.TransactionStatus   `json:"status,omitempty"`
	QrCodeID    *uint                       `json:"qr_code_id,omitempty"`
	Category    *models.TransactionCategory `json:"category,omitempty"`
	Description *string                     `json:"description,omitempty"`
}

func isValidType(t models.TransactionType) bool {
	return validTransactionTypes[t]
}

func isValidStatus(t models.TransactionStatus) bool {
	return validTransactionStatuses[t]
}

func isValidCategory(t models.TransactionCategory) bool {
	return validTransactionCategory[t]
}

func ValidateCreateTransaction(req *CreateTransactionRequest) error {

	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if req.SenderAccountID == 0 || req.ReceiverAccountID == 0 {
		return errors.New("both sender and receiver ids cannot be empty")
	}

	if req.SenderAccountID == req.ReceiverAccountID {
		return errors.New("both sender and receiver can not be same")
	}

	if req.Type != nil && *req.Type != "" && !isValidType(*req.Type) {
		return errors.New("not a valid type")
	}

	if req.Category != nil && !isValidCategory(*req.Category) {
		return errors.New("not a valid category")
	}

	return nil
}

func ValidateUpdateTransaction(req *UpdateTransactionRequest) error {

	if req.Status != nil && !isValidStatus(*req.Status) {
		return errors.New("not a valid status")
	}

	if req.Category != nil && !isValidCategory(*req.Category) {
		return errors.New("not a valid category")
	}

	return nil
}
