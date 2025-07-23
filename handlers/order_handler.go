package handlers

import (
	"errors"
	"fmt"
	"gopay-clone/models"
	"gopay-clone/services"
	"gopay-clone/utils"
	"gopay-clone/validator"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService    *services.OrderService
	merchantService *services.MerchantService
	userService     *services.UserService
	menuService     *services.MenuItemService
}

func NewOrderHandler(orderService *services.OrderService, merchantService *services.MerchantService, userService *services.UserService, menuService *services.MenuItemService) *OrderHandler {
	return &OrderHandler{orderService: orderService, merchantService: merchantService, userService: userService, menuService: menuService}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req validator.CreateOrderRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateOrder); err != nil {
		return err
	}
	loggedInUserId := utils.CLaimJwt(c)
	_, err := h.userService.GetUserById(uint(loggedInUserId))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	_, errMerchant := h.merchantService.GetMerchantByID(uint(req.MerchantID))

	if errMerchant != nil {
		fmt.Print(errMerchant)
		return utils.ValidationErrorResponse(c, errMerchant)
	}

	var totalAmount float64
	var orderItems []models.OrderItem
	for _, o := range req.Items {
		menuItem, err := h.menuService.GetMenuItemByID(o.MenuItemID)
		if err != nil || !menuItem.IsAvailable {
			return utils.ValidationErrorResponse(c, err)
		}
		itemTotal := float64(o.Quantity) * menuItem.Price
		totalAmount += itemTotal

		orderItems = append(orderItems, models.OrderItem{
			MenuItemID: o.MenuItemID,
			Quantity:   o.Quantity,
			Notes:      o.Notes,
			Price:      menuItem.Price,
		})
	}

	order := &models.Order{
		UserID:          uint(loggedInUserId),
		MerchantID:      req.MerchantID,
		DeliveryAddress: req.DeliveryAddress,
		TotalAmount:     math.Round((totalAmount+6.78)*100) / 100,
		DeliveryFee:     6.78,
	}

	if err := h.orderService.CreateOrder(order, orderItems); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	createdOrder, err := h.orderService.GetOrderByID(uint(order.ID))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Order created successfully", createdOrder)
}

func (h *OrderHandler) UpdateOrderStatus(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	var req validator.UpdateOrderStatusRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateOrderStatus); err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	if err := h.orderService.UpdateOrderStatus(uint(orderId), string(req.Status)); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Order status updated successfully", nil)
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
