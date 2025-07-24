package services

import (
	"gopay-clone/config"
	"gopay-clone/models"

	"gorm.io/gorm"
)

type OrderService struct {
	db *config.Database
}

func NewOrderService(db *config.Database) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(order *models.Order, items []models.OrderItem) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		// Create fresh OrderItem structs to avoid any ID conflicts
		for _, item := range items {
			orderItem := models.OrderItem{
				OrderID:    order.ID,
				MenuItemID: item.MenuItemID,
				Quantity:   item.Quantity,
				Price:      item.Price,
				Notes:      item.Notes,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *OrderService) GetAllOrdersByUser(id uint) ([]models.Order, error) {
	var orders []models.Order
	results := s.db.Where("user_id = ?", id).Limit(20).Find(&orders)
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

func (s *OrderService) UpdateOrderStatus(id uint, status string) error {
	return s.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.db.Delete(&models.Order{}, id).Error
}
