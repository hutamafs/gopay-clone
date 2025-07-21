package services

import (
	"errors"
	"gopay-clone/config"
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
			return err
		}
		if err := tx.First(&receiver, transaction.ReceiverAccountID).Error; err != nil {
			return err
		}
		if sender.Balance < transaction.Amount {
			return errors.New("insufficient balance")
		}

		sender.Balance -= transaction.Amount
		receiver.Balance += transaction.Amount

		if err := tx.Save(&sender).Error; err != nil {
			return err
		}
		if err := tx.Save(&receiver).Error; err != nil {
			return err
		}

		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *TransactionService) GetTransactionsByAccount(accountId uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	results := s.db.Where("sender_account_id = ?", accountId).Order("created_at DESC").Find(&transactions)
	return transactions, results.Error
}

func (s *TransactionService) GetTransactionById(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	return &transaction, s.db.
		Preload("SenderAccount").
		Preload("ReceiverAccount").
		First(&transaction, id).Error
}

func (s *TransactionService) UpdateTransaction(transactionID uint, updates map[string]any) error {
	return s.db.Model(&models.Transaction{}).Where("id = ?", transactionID).Updates(updates).Error
}
