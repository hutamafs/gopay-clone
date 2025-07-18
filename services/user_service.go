package services

import (
	"gopay-clone/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	results := s.db.Find(&users)
	return users, results.Error
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	var user models.User
	return &user, s.db.First(&user, id).Error
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.db.Save(&user).Error
}

func (s *UserService) DeleteUser(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}
