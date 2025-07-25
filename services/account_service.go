package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
	"gopay-clone/models"

	"gorm.io/gorm"
)

type AccountService struct {
	db *config.Database
}

func NewAccountService(db *config.Database) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	if err := s.db.Create(account).Error; err != nil {
		return apperrors.ErrAccountCreateFailed
	}
	return nil
}

func (s *AccountService) GetAccountsByUser(userId uint) ([]models.Account, error) {
	var accounts []models.Account
	if err := s.db.Where("user_id = ?", userId).Find(&accounts).Error; err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch user accounts")
	}
	return accounts, nil
}

func (s *AccountService) GetAccountById(id uint) (*models.Account, error) {
	var account models.Account
	err := s.db.
		Preload("SentTransactions").
		Preload("ReceivedTransactions").
		First(&account, id).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrAccountNotFound
		}
		return nil, apperrors.NewInternalError("Failed to fetch account")
	}
	return &account, nil
}

func (s *AccountService) GetMainBalanceAccount(userID uint) (*models.Account, error) {
	var account models.Account
	err := s.db.Where("user_id = ? AND account_type = ?", userID, "main_balance").First(&account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrAccountNotFound
		}
		return nil, apperrors.NewInternalError("Failed to fetch main balance account")
	}
	return &account, nil
}

func (s *AccountService) GetBalanceByAccountId(accountId uint) (*float64, error) {
	var account models.Account
	err := s.db.First(&account, accountId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrAccountNotFound
		}
		return nil, apperrors.NewInternalError("Failed to fetch account")
	}
	return &account.Balance, nil
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	if err := s.db.Model(&account).Select("name").Updates(account).Error; err != nil {
		return apperrors.ErrAccountUpdateFailed
	}
	return nil
}
