package services

import (
	"gopay-clone/config"
	"gopay-clone/models"
)

type UserService struct {
	db *config.Database
}

func NewUserService(db *config.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}
	defaultAccounts := []models.Account{
		{
			Name:        "main wallet",
			AccountType: "main",
			UserId:      user.ID,
		},
		{
			Name:        "points",
			AccountType: "gopay_points",
			UserId:      user.ID,
		},
	}
	for _, account := range defaultAccounts {
		if err := s.db.Create(&account).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	results := s.db.Find(&users)
	return users, results.Error
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	var user models.User
	result := s.db.
		Preload("Accounts").
		Preload("Contacts").
		First(&user, id)
	return &user, result.Error
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.db.Model(&user).Select("name", "password", "updated_at").Updates(user).Error
}

func (s *UserService) DeleteUser(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}
