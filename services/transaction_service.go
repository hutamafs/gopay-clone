package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
	"gopay-clone/models"

	"gorm.io/gorm"
)

type TransactionService struct {
	db *config.Database
}

func NewTransactionService(db *config.Database) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	var sender models.Account
	var receiver models.Account
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&sender, transaction.SenderAccountID).Error; err != nil {
			return apperrors.ErrAccountNotFound
		}
		if err := tx.First(&receiver, transaction.ReceiverAccountID).Error; err != nil {
			return apperrors.ErrAccountNotFound
		}
		if sender.Balance < transaction.Amount {
			return apperrors.ErrInsufficientBalance
		}
		if sender.ID == receiver.ID {
			return apperrors.ErrSameAccount
		}

		sender.Balance -= transaction.Amount
		receiver.Balance += transaction.Amount

		if err := tx.Save(&sender).Error; err != nil {
			return apperrors.ErrDatabaseError
		}
		if err := tx.Save(&receiver).Error; err != nil {
			return apperrors.ErrDatabaseError
		}

		if err := tx.Create(transaction).Error; err != nil {
			return apperrors.ErrTransactionFailed
		}

		return nil
	})
}

func (s *TransactionService) GetTransactionsByAccount(accountId uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := s.db.Where("sender_account_id = ?", accountId).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, apperrors.ErrTransactionNotFound
	}
	return transactions, nil
}

func (s *TransactionService) GetTransactionById(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := s.db.
		Preload("SenderAccount").
		Preload("ReceiverAccount").
		First(&transaction, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrTransactionNotFound
		}
		return nil, apperrors.ErrDatabaseError
	}
	return &transaction, nil
}

func (s *TransactionService) UpdateTransaction(transactionID uint, updates map[string]any) error {
	result := s.db.Model(&models.Transaction{}).Where("id = ?", transactionID).Updates(updates)
	if result.Error != nil {
		return apperrors.ErrTransactionFailed
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrTransactionNotFound
	}
	return nil
}

func (s *TransactionService) UpdateTransactionWhenFoodOrderCompleted(orderId uint) error {
	result := s.db.Model(&models.Transaction{}).Where("service_type = ? AND service_id = ?", "food", orderId).Update("status", "completed")
	if result.Error != nil {
		return apperrors.ErrTransactionFailed
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrTransactionNotFound
	}
	return nil
}
