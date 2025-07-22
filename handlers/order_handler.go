package handlers

import (
	"errors"
	"gopay-clone/models"
	"gopay-clone/services"
	"gopay-clone/utils"
	"gopay-clone/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService    *services.OrderService
	merchantService *services.MerchantService
	userService     *services.UserService
}

func NewOrderHandler(orderService *services.OrderService, merchantService *services.MerchantService, userService *services.UserService) *OrderHandler {
	return &OrderHandler{orderService: orderService, merchantService: merchantService, userService: userService}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req validator.CreateOrderRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateOrder); err != nil {
		return err
	}
	_, err := h.userService.GetUserById(req.UserID)
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	_, errMerchant := h.merchantService.GetMerchantByID(req.MerchantID)
	if errMerchant != nil {
		return utils.ValidationErrorResponse(c, errMerchant)
	}
	var items []validator.CreateOrderItemRequest
	for _, o := range req.Items {
		items = append(items, validator.CreateOrderItemRequest{
			MenuItemID: o.MenuItemID,
			Quantity:   o.Quantity,
			Notes:      o.Notes,
			Price:      o.Price,
		})
	}

	order := &models.Order{
		UserID:          req.UserID,
		MerchantID:      req.MerchantID,
		DeliveryAddress: req.DeliveryAddress,
	}

	if err := h.orderService.CreateOrder(order, items); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Order created successfully", order)
}

func (h *OrderHandler) GetOrderByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	order, err := h.orderService.GetOrderByID(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Order detail fetched successfully", order)
}

func (h *OrderHandler) UpdateMerchantByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	merchant, err := h.merchantService.GetMerchantByID(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	loggedInUserId := utils.CLaimJwt(c)

	if merchant.UserId != uint(loggedInUserId) {
		return utils.ForbiddenResponse(c, errors.New("unauthorized access"))
	}
	var req validator.UpdateMerchantRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateMerchant); err != nil {
		return err
	}

	updatedMerchant := map[string]any{
		"UserId":          merchant.UserId,
		"Location":        req.Location,
		"MerchantName":    req.MerchantName,
		"Description":     req.Description,
		"MerchantPhone":   req.MerchantPhone,
		"Category":        req.Category,
		"OpenHour":        req.OpenHour,
		"ClosedHour":      req.ClosedHour,
		"MerchantLogoURL": req.MerchantLogoURL,
	}

	if err := h.merchantService.UpdateMerchant(uint(id), updatedMerchant); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Merchant profile updated successfully", merchant)
}
