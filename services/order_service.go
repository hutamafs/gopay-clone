package services

import (
	"gopay-clone/config"
	"gopay-clone/models"
	"gopay-clone/validator"
)

type OrderService struct {
	db *config.Database
}

func NewOrderService(db *config.Database) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(order *models.Order, items []validator.CreateOrderItemRequest) error {
	if err := s.db.Create(order).Error; err != nil {
		return err
	}
	for _, i := range items {
		oi := models.OrderItem{
			OrderID:    order.ID,
			MenuItemID: i.MenuItemID,
			Quantity:   i.Quantity,
			Price:      i.Price,
			Notes:      i.Notes,
		}
		if err := s.db.Create(&oi).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) GetAllOrdersByUser(id uint) ([]models.Order, error) {
	var orders []models.Order
	results := s.db.Find(&orders).Where("user_id = ?", id).Limit(20)
	return orders, results.Error
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	result := s.db.Preload("User").
		Preload("Merchant").
		Preload("Items").
		First(&order, id)
	return &order, result.Error
}

func (s *OrderService) GetMerchantByUserID(id uint) (*models.MerchantProfile, error) {
	var merchant models.MerchantProfile
	result := s.db.
		Where("user_id = ?", id).
		First(&merchant)
	return &merchant, result.Error
}

func (s *OrderService) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

func (s *OrderService) UpdateMerchant(id uint, merchant map[string]any) error {
	return s.db.Model(&models.MerchantProfile{}).Where("id = ?", id).Updates(merchant).Error
}

func (s *OrderService) DeleteMerchant(id uint) error {
	return s.db.Delete(&models.MerchantProfile{}, id).Error
}
