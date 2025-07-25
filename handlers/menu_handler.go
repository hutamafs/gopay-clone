package handlers

import (
	"errors"
	"fmt"
	"gopay-clone/models"
	"gopay-clone/services"
	"gopay-clone/utils"
	"gopay-clone/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenuHandler struct {
	menuService     *services.MenuItemService
	merchantService *services.MerchantService
}

func NewMenuHandler(menuService *services.MenuItemService, merchantService *services.MerchantService) *MenuHandler {
	return &MenuHandler{menuService: menuService, merchantService: merchantService}
}

func (h *MenuHandler) CreateMenu(c echo.Context) error {
	var req validator.CreateMenuRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateMenu); err != nil {
		return err
	}
	user_id, user_role := utils.GetIDAndRoleFromJWT(c)
	m, err := h.merchantService.GetMerchantByUserID(uint(user_id))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	if user_role != "merchant" {
		return utils.ValidationErrorResponse(c, errors.New("only merchant can create menu"))
	}
	var category models.MenuCategory
	if req.Category != nil {
		category = models.MenuCategory(*req.Category)
	}

	menu := &models.MenuItem{
		Name:         req.Name,
		MerchantId:   m.ID,
		Description:  req.Description,
		Price:        req.Price,
		MenuImageURL: req.MenuImageURL,
		Category:     models.MenuCategory(category),
	}

	if err := h.menuService.CreateMenu(menu); err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Menu created successfully", menu)
}

func (h *MenuHandler) GetAllMenus(c echo.Context) error {
	merchant_id, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	menus, err := h.menuService.GetAllMenusFromMerchant(uint(merchant_id))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "All Menus fetched successfully", menus)
}

func (h *MenuHandler) GetMenuByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("menu_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	menu, err := h.menuService.GetMenuItemByID(uint(id))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Menu detail fetched successfully", menu)
}

func (h *MenuHandler) UpdateMenuItem(c echo.Context) error {
	merchantId, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	menuId, err := strconv.Atoi(c.Param("menu_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	// Verify ownership logged in id == the merchant user.id
	merchant, err := h.merchantService.GetMerchantByID(uint(merchantId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	loggedInUserId := utils.CLaimJwt(c)
	if merchant.UserId != uint(loggedInUserId) {
		return utils.ValidationErrorResponse(c, errors.New("unauthorized: merchant can only update their own menu items"))
	}

	// verify menu item belongs to merchant
	menuItem, err := h.menuService.GetMenuItemByID(uint(menuId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	if menuItem.MerchantId != uint(merchantId) {
		return utils.ValidationErrorResponse(c, errors.New("unauthorized: menu item does not belong to this merchant"))
	}

	var req validator.UpdateMenuItemRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateMenuItem); err != nil {
		return err
	}
	var category models.MenuCategory
	if req.Category != nil {
		category = models.MenuCategory(*req.Category)
	}

	updates := map[string]any{
		"name":           req.Name,
		"description":    req.Description,
		"price":          req.Price,
		"category":       category,
		"menu_image_url": req.MenuImageURL,
	}

	// only update availability if provided
	if req.IsAvailable != nil {
		updates["is_available"] = *req.IsAvailable
	}

	if err := h.menuService.UpdateMenuItem(uint(menuId), updates); err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	updatedMenuItem, _ := h.menuService.GetMenuItemByID(uint(menuId))
	return utils.SuccessResponse(c, http.StatusOK, "Menu item updated successfully", updatedMenuItem)
}

func (h *MenuHandler) DeleteMenuItem(c echo.Context) error {
	merchantId, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	menuId, err := strconv.Atoi(c.Param("menu_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	// Verify ownership logged in id == the merchant user.id
	merchant, err := h.merchantService.GetMerchantByID(uint(merchantId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	loggedInUserId := utils.CLaimJwt(c)
	if merchant.UserId != uint(loggedInUserId) {
		return utils.ValidationErrorResponse(c, errors.New("unauthorized: merchant can only delete their own menu items"))
	}

	// verify menu item belongs to merchant
	menuItem, err := h.menuService.GetMenuItemByID(uint(menuId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	if menuItem.MerchantId != uint(merchantId) {
		return utils.ValidationErrorResponse(c, errors.New("unauthorized: menu item does not belong to this merchant"))
	}

	if err := h.menuService.DeleteMenuItem(uint(menuId)); err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Deleted Menu Item", nil)
}

func (h *MenuHandler) GetAllMenusByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	menus, err := h.menuService.GetMenuByCategory(category)
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	str := fmt.Sprintf("menus with category %v has been fetched successfully", category)
	return utils.SuccessResponse(c, http.StatusOK, str, menus)
}
