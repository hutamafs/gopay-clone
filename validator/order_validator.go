package validator

import (
	"errors"
	"gopay-clone/models"
	"strings"
)

var validOrderStatuses = map[models.OrderStatus]bool{
	models.OrderPending:   true,
	models.OrderConfirmed: true,
	models.OrderPreparing: true,
	models.OrderReady:     true,
	models.OrderDelivery:  true,
	models.OrderCompleted: true,
	models.OrderCancelled: true,
}

type CreateOrderRequest struct {
	MerchantID      uint                     `json:"merchant_id" validate:"required"`
	DeliveryAddress string                   `json:"delivery_address" validate:"required"`
	Items           []CreateOrderItemRequest `json:"order_items" validate:"required,min=1"`
}

type CreateOrderItemRequest struct {
	MenuItemID uint   `json:"menu_item_id" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required,min=1"`
	Notes      string `json:"notes"`
}

type UpdateOrderStatusRequest struct {
	Status models.OrderStatus `json:"status" validate:"required"`
}

func ValidateCreateOrder(req *CreateOrderRequest) error {
	if req.MerchantID == 0 {
		return errors.New("merchant id cannot be empty")
	}
	if strings.TrimSpace(req.DeliveryAddress) == "" {
		return errors.New("delivery address cannot be empty")
	}
	if len(req.Items) < 1 {
		return errors.New("need order items")
	}

	return nil
}

func isValidOrderStatus(t models.OrderStatus) bool {
	return validOrderStatuses[t]
}

func ValidateUpdateOrderStatus(req *UpdateOrderStatusRequest) error {
	if !isValidOrderStatus(models.OrderStatus(req.Status)) {
		return errors.New("not a valid order status")
	}
	validStatuses := []string{"pending", "confirmed", "cooking", "ready", "delivery", "completed", "cancelled"}

	for _, validStatus := range validStatuses {
		if string(req.Status) == validStatus {
			return nil
		}
	}

	return nil
}
