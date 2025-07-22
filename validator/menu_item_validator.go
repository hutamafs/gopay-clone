package validator

import (
	"errors"
	"gopay-clone/models"
)

var validMenuCategory = map[models.MenuCategory]bool{
	models.MainCourse: true,
	models.Appetizer:  true,
	models.Dessert:    true,
	models.Beverage:   true,
	models.Snack:      true,
	models.Drink:      true,
}

type CreateMenuRequest struct {
	MerchantId   uint                 `json:"merchant_id" validate:"required"`
	Name         string               `json:"name" validate:"required"`
	Description  string               `json:"description"`
	Price        float64              `json:"price" validate:"required"`
	MenuImageURL string               `json:"merchant_logo_url"`
	Category     *models.MenuCategory `json:"category"`
}

// type UpdateMerchantRequest struct {
// 	Location        string `json:"location" validate:"required"`
// 	MerchantName    string `json:"merchant_name" validate:"required"`
// 	Description     string `json:"description"`
// 	MerchantPhone   string `json:"merchant_phone"`
// 	Category        string `json:"category"`
// 	OpenHour        string `json:"open_hour" validate:"required"`
// 	ClosedHour      string `json:"closed_hour" validate:"required"`
// 	MerchantLogoURL string `json:"merchant_logo_url"`
// }

func ValidateCreateMenu(req *CreateMenuRequest) error {
	if err := validateEmptyString(req.Name, "menu name"); err != nil {
		return err
	}
	if req.Price <= 0 {
		return errors.New("price can not be 0")
	}
	if req.Category != nil {
		if !isValidMenuCategory(*req.Category) {
			return errors.New("menu category is not valid")
		}
	}
	return nil
}

func isValidMenuCategory(t models.MenuCategory) bool {
	return validMenuCategory[t]
}
