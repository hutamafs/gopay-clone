package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
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
	var existingMerchant models.MerchantProfile
	if err := s.db.Where("user_id = ?", merchant.UserId).First(&existingMerchant).Error; err == nil {
		return apperrors.ErrMerchantExists
	}
	if err := s.db.Create(merchant).Error; err != nil {
		return apperrors.ErrMerchantCreateFailed
	}
	return nil
}

func (s *MerchantService) GetAllMerchants() ([]models.MerchantProfile, error) {
	var merchants []models.MerchantProfile
	if err := s.db.Find(&merchants).Error; err != nil {
		return nil, apperrors.ErrDatabaseError
	}
	return merchants, nil
}

func (s *MerchantService) GetMerchantByID(id uint) (*models.MerchantProfile, error) {
	var merchant models.MerchantProfile
	if err := s.db.Preload("Menu").
		First(&merchant, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrMerchantNotFound
		}
		return nil, apperrors.ErrDatabaseError
	}
	return &merchant, nil
}

func (s *MerchantService) GetMerchantByUserID(id uint) (*models.MerchantProfile, error) {
	var merchant models.MerchantProfile
	if err := s.db.
		Where("user_id = ?", id).
		First(&merchant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrMerchantNotFound
		}
		return nil, apperrors.ErrDatabaseError
	}
	return &merchant, nil
}

func (s *MerchantService) UpdateMerchant(id uint, merchant map[string]any) error {
	result := s.db.Model(&models.MerchantProfile{}).Where("id = ?", id).Updates(merchant)
	if result.Error != nil {
		return apperrors.ErrMerchantUpdateFailed
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrMerchantNotFound
	}
	return nil
}

func (s *MerchantService) DeleteMerchant(id uint) error {
	result := s.db.Delete(&models.MerchantProfile{}, id)
	if result.Error != nil {
		return apperrors.ErrMerchantDeleteFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrMerchantNotFound
	}

	return nil
}
