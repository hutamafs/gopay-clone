package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
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
			return apperrors.ErrOrderCreateFailed
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
				return apperrors.ErrOrderCreateFailed
			}
		}
		return nil
	})
}

func (s *OrderService) GetAllOrdersByUser(id uint) ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Where("user_id = ?", id).Limit(20).Find(&orders).Error; err != nil {
		return nil, apperrors.ErrDatabaseError
	}
	return orders, nil
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := s.db.Preload("User").
		Preload("Merchant").
		Preload("Items").
		First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrOrderNotFound
		}
		return nil, apperrors.ErrDatabaseError
	}
	return &order, nil
}

func (s *OrderService) UpdateOrderStatus(id uint, status string) error {
	result := s.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return apperrors.ErrOrderStatusUpdateFailed
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrOrderNotFound
	}
	return nil
}

func (s *OrderService) DeleteOrder(id uint) error {
	result := s.db.Delete(&models.Order{}, id)
	if result.Error != nil {
		return apperrors.ErrOrderDeleteFailed
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrOrderNotFound
	}
	return nil
}
