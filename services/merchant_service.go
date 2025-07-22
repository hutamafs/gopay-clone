package services

import (
	"gopay-clone/config"
	"gopay-clone/models"

	"gorm.io/gorm"
)

type MerchantService struct {
	db *config.Database
}

func NewMerchantService(db *config.Database) *MerchantService {
	return &MerchantService{db: db}
}

func (s *MerchantService) CreateMerchant(merchant *models.MerchantProfile) error {
	return s.db.Create(merchant).Error
}

func (s *MerchantService) GetAllMerchants() ([]models.MerchantProfile, error) {
	var merchants []models.MerchantProfile
	results := s.db.Find(&merchants)
	return merchants, results.Error
}

func (s *MerchantService) GetMerchantByID(id uint) (*models.MerchantProfile, error) {
	var merchant models.MerchantProfile
	result := s.db.Preload("Menu", func(db *gorm.DB) *gorm.DB {
		return db.
			Order("rating DESC").
			Limit(10)
	}).
		First(&merchant, id)
	return &merchant, result.Error
}

func (s *MerchantService) GetMerchantByUserID(id uint) (*models.MerchantProfile, error) {
	var merchant models.MerchantProfile
	result := s.db.
		Where("user_id = ?", id).
		First(&merchant)
	return &merchant, result.Error
}

func (s *MerchantService) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

func (s *MerchantService) UpdateMerchant(id uint, merchant map[string]any) error {
	return s.db.Model(&models.MerchantProfile{}).Where("id = ?", id).Updates(merchant).Error
}

func (s *MerchantService) DeleteMerchant(id uint) error {
	return s.db.Delete(&models.MerchantProfile{}, id).Error
}
