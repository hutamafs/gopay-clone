package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
	"gopay-clone/models"

	"gorm.io/gorm"
)

type MenuItemService struct {
	db *config.Database
}

func NewMenuItemService(db *config.Database) *MenuItemService {
	return &MenuItemService{db: db}
}

func (s *MenuItemService) CreateMenu(menu *models.MenuItem) error {
	if err := s.db.Create(menu).Error; err != nil {
		return apperrors.ErrMenuCreateFailed
	}
	return nil
}

func (s *MenuItemService) GetAllMenusFromMerchant(merchantId uint) ([]models.MenuItem, error) {
	var menuItems []models.MenuItem
	if err := s.db.Where("merchant_id = ?", merchantId).
		Order("category ASC, name ASC").
		Find(&menuItems).Error; err != nil {
		return nil, apperrors.ErrDatabaseError
	}
	return menuItems, nil
}

func (s *MenuItemService) GetMenuItemByID(id uint) (*models.MenuItem, error) {
	var menuItem models.MenuItem
	if err := s.db.Preload("Merchant").First(&menuItem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrMenuNotFound
		}
		return nil, apperrors.ErrDatabaseError
	}
	return &menuItem, nil
}

func (s *MenuItemService) UpdateMenuItem(id uint, menu map[string]any) error {
	result := s.db.Model(&models.MenuItem{}).Where("id = ?", id).Updates(menu)
	if result.Error != nil {
		return apperrors.ErrMenuUpdateFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrMenuNotFound
	}

	return nil
}

func (s *MenuItemService) DeleteMenuItem(id uint) error {
	result := s.db.Delete(&models.MenuItem{}, id)
	if result.Error != nil {
		return apperrors.ErrMenuDeleteFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrMenuNotFound
	}

	return nil
}

func (s *MenuItemService) GetMenuByCategory(category string) ([]models.MenuItem, error) {
	var menuItems []models.MenuItem
	if err := s.db.Where("category = ?", category).
		Order("name ASC").
		Find(&menuItems).Error; err != nil {
		return nil, apperrors.ErrDatabaseError
	}
	return menuItems, nil
}
