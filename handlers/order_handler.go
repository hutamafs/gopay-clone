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
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService       *services.OrderService
	merchantService    *services.MerchantService
	userService        *services.UserService
	menuService        *services.MenuItemService
	accountService     *services.AccountService
	transactionService *services.TransactionService
	driverService      *services.DriverService
}

func NewOrderHandler(
	orderService *services.OrderService,
	merchantService *services.MerchantService, userService *services.UserService,
	menuService *services.MenuItemService,
	accountService *services.AccountService,
	transactionService *services.TransactionService,
	driverService *services.DriverService,
) *OrderHandler {
	return &OrderHandler{
		orderService:       orderService,
		merchantService:    merchantService,
		userService:        userService,
		menuService:        menuService,
		accountService:     accountService,
		transactionService: transactionService,
		driverService:      driverService,
	}
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
	totalAmount = math.Round((totalAmount+6.78)*100) / 100
	userAccount, er := h.accountService.GetMainBalanceAccount(uint(loggedInUserId))
	if er != nil || userAccount == nil {
		return utils.ValidationErrorResponse(c, errors.New("user account not found"))
	}
	merchantAccount, e := h.accountService.GetMainBalanceAccount(uint(req.MerchantID))
	if e != nil || merchantAccount == nil {
		return utils.ValidationErrorResponse(c, errors.New("merchant account not found"))
	}
	// for now, point system will now be used yet
	if userAccount.Balance < totalAmount {
		return utils.ValidationErrorResponse(c, errors.New("insufficient balance"))
	}

	// create the order
	order := &models.Order{
		UserID:          uint(loggedInUserId),
		MerchantID:      req.MerchantID,
		DeliveryAddress: req.DeliveryAddress,
		TotalAmount:     totalAmount,
		DeliveryFee:     6.78,
	}

	// get the driver
	availableDrivers, err := h.driverService.GetAvailableDrivers()
	if err != nil || len(availableDrivers) == 0 {
		return utils.ValidationErrorResponse(c, errors.New("no available driver"))
	}
	selectedDriver := availableDrivers[0]
	order.DriverID = &selectedDriver.ID

	if err := h.driverService.UpdateDriverStatus(selectedDriver.ID, "sending"); err != nil {
		return err
	}

	if err := h.orderService.CreateOrder(order, orderItems); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	transaction := &models.Transaction{
		Amount:            totalAmount,
		SenderAccountID:   userAccount.ID,
		ReceiverAccountID: merchantAccount.ID,
		Category:          "food",
		Type:              "payment",
		Status:            "pending",
		ServiceType:       "food",
		ServiceID:         &order.ID,
	}

	if err := h.transactionService.CreateTransaction(transaction); err != nil {
		h.orderService.DeleteOrder(order.ID)
		return utils.ValidationErrorResponse(c, err)
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

	loggedInUserId := utils.CLaimJwt(c)
	order, err := h.orderService.GetOrderByID(uint(orderId))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}

	// check if the logged in id can update it
	if err := h.validateStatusUpdate(order, uint(loggedInUserId), string(req.Status)); err != nil {
		return utils.ForbiddenResponse(c, errors.New("unauthorized to update the status"))
	}

	// do the actual update
	if err := h.orderService.UpdateOrderStatus(uint(orderId), string(req.Status)); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	// update the transaction as well if the status is completed
	if req.Status == models.OrderCompleted {
		if err := h.transactionService.UpdateTransactionWhenFoodOrderCompleted(order.ID); err != nil {
			return utils.InternalErrorResponse(c, err)
		}
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

func (h *OrderHandler) validateStatusUpdate(order *models.Order, userID uint, newStatus string) error {
	currentStatus := string(order.Status)

	// check if user is merchant (owns the merchant of this order)
	merchant, err := h.merchantService.GetMerchantByID(order.MerchantID)
	if err == nil && merchant.UserId == userID {
		return h.validateMerchantStatusUpdate(currentStatus, newStatus)
	}

	// check if user is the driver
	if order.DriverID != nil {
		driver, err := h.driverService.GetDriverByID(*order.DriverID)
		if err == nil && driver.UserId == userID {
			return h.validateDriverStatusUpdate(currentStatus, newStatus)
		}
	}

	// check if user is the customer who order the food
	if order.UserID == userID {
		return h.validateCustomerStatusUpdate(currentStatus, newStatus)
	}

	return errors.New("unauthorized to update this order")
}

func (h *OrderHandler) validateMerchantStatusUpdate(current, new string) error {
	validTransitions := map[string][]string{
		"pending":   {"confirmed", "cancelled"},
		"confirmed": {"cooking"},
		"cooking":   {"ready"},
	}

	// validnext -> get the slices
	// valid transition[current] -> let's say current is pending, it will get the slices {confirmed cancelled}
	if validNext, exists := validTransitions[current]; exists {
		if slices.Contains(validNext, new) {
			return nil
		}
	}

	return errors.New("invalid status transition for merchant")
}

func (h *OrderHandler) validateDriverStatusUpdate(current, new string) error {
	validTransitions := map[string][]string{
		"ready":    {"delivery"},
		"delivery": {"completed"},
	}

	if validNext, exists := validTransitions[current]; exists {
		if slices.Contains(validNext, new) {
			return nil
		}
	}

	return errors.New("invalid status transition for driver")
}

func (h *OrderHandler) validateCustomerStatusUpdate(current, new string) error {
	// customers can only cancel pending orders
	if current == "pending" && new == "cancelled" {
		return nil
	}

	return errors.New("customers can only cancel pending orders")
}
