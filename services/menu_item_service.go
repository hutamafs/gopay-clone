package services

import (
	"gopay-clone/config"
	"gopay-clone/models"
)

type MenuItemService struct {
	db *config.Database
}

func NewMenuItemService(db *config.Database) *MenuItemService {
	return &MenuItemService{db: db}
}

func (s *MenuItemService) CreateMenu(menu *models.MenuItem) error {
	return s.db.Create(menu).Error
}

func (s *MenuItemService) GetAllMenusFromMerchant(merchantId uint) ([]models.MenuItem, error) {
	var menuItems []models.MenuItem
	result := s.db.Where("merchant_id = ?", merchantId).
		Order("category ASC, name ASC").
		Find(&menuItems)
	return menuItems, result.Error
}

func (s *MenuItemService) GetMenuItemByID(id uint) (*models.MenuItem, error) {
	var menuItem models.MenuItem
	result := s.db.Preload("Merchant").First(&menuItem, id)
	return &menuItem, result.Error
}

func (s *MenuItemService) UpdateMenuItem(id uint, menu map[string]any) error {
	return s.db.Model(&models.MenuItem{}).Where("id = ?", id).Updates(menu).Error
}

func (s *MenuItemService) DeleteMenuItem(id uint) error {
	return s.db.Delete(&models.MenuItem{}, id).Error
}

func (s *MenuItemService) GetMenuByCategory(category string) ([]models.MenuItem, error) {
	var menuItems []models.MenuItem
	result := s.db.Where("category = ?", category).
		Order("name ASC").
		Find(&menuItems)
	return menuItems, result.Error
}
