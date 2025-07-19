package services

import (
	"gopay-clone/config"
	"gopay-clone/models"
)

type AccountService struct {
	db *config.Database
}

func NewAccountService(db *config.Database) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.db.Create(account).Error
}

func (s *AccountService) GetAccountsByUser(userId uint) ([]models.Account, error) {
	var accounts []models.Account
	results := s.db.Where("user_id = ?", userId).Find(&accounts)
	return accounts, results.Error
}

func (s *AccountService) GetAccountById(id uint) (*models.Account, error) {
	var account models.Account
	return &account, s.db.First(&account, id).Error
}

func (s *AccountService) GetBalanceByAccountId(accountId uint) (*float64, error) {
	var account models.Account
	return &account.Balance, s.db.First(&account, accountId).Error
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.db.Save(&account).Error
}
