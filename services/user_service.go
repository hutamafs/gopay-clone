package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
	"gopay-clone/models"
	"gopay-clone/utils"

	"gorm.io/gorm"
)

type UserService struct {
	db *config.Database
}

func NewUserService(db *config.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Check if user already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return apperrors.ErrUserExists
	}

	if err := s.db.Create(user).Error; err != nil {
		return apperrors.NewInternalError("Failed to create user")
	}

	// Create default accounts
	defaultAccounts := []models.Account{
		{
			Name:        "main wallet",
			AccountType: models.MainBalance,
			UserId:      user.ID,
		},
		{
			Name:        "points",
			AccountType: models.Points,
			UserId:      user.ID,
		},
	}

	for _, account := range defaultAccounts {
		if err := s.db.Create(&account).Error; err != nil {
			return apperrors.ErrAccountCreateFailed
		}
	}

	return nil
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch users")
	}
	return users, nil
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	var user models.User
	err := s.db.
		Preload("Accounts").
		Preload("Contacts").
		First(&user, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, apperrors.NewInternalError("Failed to fetch user")
	}

	return &user, nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	if err := s.db.Save(user).Error; err != nil {
		return apperrors.NewInternalError("Failed to update user")
	}
	return nil
}

func (s *UserService) DeleteUser(id uint) error {
	result := s.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return apperrors.NewInternalError("Failed to delete user")
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrUserNotFound
	}
	return nil
}

func (s *UserService) Login(user *models.LoggedinUser) (string, error) {
	var foundUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", apperrors.ErrEmailNotFound
		}
		return "", apperrors.NewInternalError("Database error during login")
	}

	if !utils.CheckPassword(foundUser.Password, user.Password) {
		return "", apperrors.ErrInvalidPassword
	}

	tokenString, err := utils.CreateToken(foundUser)
	if err != nil {
		return "", apperrors.ErrTokenCreation
	}

	return tokenString, nil
}
