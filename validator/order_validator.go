package validator

import (
	"errors"
	"strings"
)

type CreateOrderRequest struct {
	UserID          uint                     `json:"user_id" validate:"required"`
	MerchantID      uint                     `json:"merchant_id" validate:"required"`
	DeliveryAddress string                   `json:"delivery_address" validate:"required"`
	Items           []CreateOrderItemRequest `json:"order_items" validate:"required,min=1"`
}

type CreateOrderItemRequest struct {
	MenuItemID uint    `json:"menu_item_id" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required,min=1"`
	Price      float64 `json:"price" validate:"required"`
	Notes      string  `json:"notes"`
}

func ValidateCreateOrder(req *CreateOrderRequest) error {
	if req.UserID == 0 {
		return errors.New("user id cannot be empty")
	}
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

func ValidateUpdateOrder(req *UpdateAccountRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("account name cannot be empty")
	}

	return nil
}
