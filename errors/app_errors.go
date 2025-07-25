package errors

import "net/http"

// AppError represents a custom application error
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	HTTPStatus int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

// user-related errors
var (
	ErrDatabaseError   = &AppError{"DATABASE_ERROR", "Database operation failed", "internal", http.StatusInternalServerError}
	ErrUserNotFound    = &AppError{"USER_NOT_FOUND", "User not found", "not_found", http.StatusNotFound}
	ErrUserExists      = &AppError{"USER_EXISTS", "User already exists", "conflict", http.StatusConflict}
	ErrInvalidPassword = &AppError{"INVALID_PASSWORD", "Invalid password", "unauthorized", http.StatusUnauthorized}
	ErrEmailNotFound   = &AppError{"EMAIL_NOT_FOUND", "Email not found", "unauthorized", http.StatusUnauthorized}
	ErrTokenCreation   = &AppError{"TOKEN_CREATION_FAILED", "Failed to create authentication token", "internal", http.StatusInternalServerError}
	ErrUserCreation    = &AppError{"USER_CREATION_FAILED", "Failed to create user", "internal", http.StatusInternalServerError}
)

// driver-related errors
var (
	ErrDriverNotFound            = &AppError{"DRIVER_NOT_FOUND", "Driver not found", "not_found", http.StatusNotFound}
	ErrDriverExists              = &AppError{"DRIVER_EXISTS", "Driver already exists", "conflict", http.StatusConflict}
	ErrDriverUnavailable         = &AppError{"DRIVER_UNAVAILABLE", "No Available drivers", "not_found", http.StatusNotFound}
	ErrDriverCreation            = &AppError{"DRIVER_CREATION_FAILED", "Failed to create driver", "internal", http.StatusInternalServerError}
	ErrDriverProfileUpdateFailed = &AppError{"DRIVER_PROFILE_UPDATE_FAILED", "Failed to update driver profile", "internal", http.StatusInternalServerError}
	ErrDriverStatusUpdateFailed  = &AppError{"DRIVER_STATUS_UPDATE_FAILED", "Failed to update driver status", "internal", http.StatusInternalServerError}
	ErrDriverDeleteFailed        = &AppError{"DRIVER_DELETE_FAILED", "Failed to delete driver ", "internal", http.StatusInternalServerError}
)

// account-related errors
var (
	ErrAccountNotFound     = &AppError{"ACCOUNT_NOT_FOUND", "Account not found", "not_found", http.StatusNotFound}
	ErrInsufficientBalance = &AppError{"INSUFFICIENT_BALANCE", "Insufficient balance", "validation", http.StatusBadRequest}
	ErrInvalidAccountType  = &AppError{"INVALID_ACCOUNT_TYPE", "Invalid account type", "validation", http.StatusBadRequest}
	ErrAccountCreateFailed = &AppError{"ACCOUNT_CREATE_FAILED", "Failed to create account", "internal", http.StatusInternalServerError}
	ErrAccountUpdateFailed = &AppError{"ACCOUNT_UPDATE_FAILED", "Failed to update account", "internal", http.StatusInternalServerError}
)

// transaction-related errors
var (
	ErrTransactionNotFound = &AppError{"TRANSACTION_NOT_FOUND", "Transaction not found", "not_found", http.StatusNotFound}
	ErrInvalidAmount       = &AppError{"INVALID_AMOUNT", "Amount must be greater than zero", "validation", http.StatusBadRequest}
	ErrSameAccount         = &AppError{"SAME_ACCOUNT_TRANSFER", "Cannot transfer to the same account", "validation", http.StatusBadRequest}
	ErrTransactionFailed   = &AppError{"TRANSACTION_FAILED", "Transaction processing failed", "internal", http.StatusInternalServerError}
)

// QR Code-related errors
var (
	ErrQRNotFound     = &AppError{"QR_NOT_FOUND", "QR code not found", "not_found", http.StatusNotFound}
	ErrQRExpired      = &AppError{"QR_EXPIRED", "QR code has expired", "validation", http.StatusBadRequest}
	ErrQRAlreadyUsed  = &AppError{"QR_ALREADY_USED", "QR code has already been used", "validation", http.StatusBadRequest}
	ErrQRCreateFailed = &AppError{"QR_CREATE_FAILED", "Failed to create QR code", "internal", http.StatusInternalServerError}
)

// order-related errors
var (
	ErrOrderNotFound           = &AppError{"ORDER_NOT_FOUND", "Order not found", "not_found", http.StatusNotFound}
	ErrOrderCreateFailed       = &AppError{"ORDER_CREATE_FAILED", "Failed to create order", "internal", http.StatusInternalServerError}
	ErrOrderStatusUpdateFailed = &AppError{"ORDER_STATUS_UPDATE_FAILED", "Failed to update order status", "internal", http.StatusInternalServerError}
	ErrOrderDeleteFailed       = &AppError{"ORDER_DELETE_FAILED", "Failed to delete order", "internal", http.StatusInternalServerError}
)

// menu-related errors
var (
	ErrMenuNotFound     = &AppError{"MENU_NOT_FOUND", "Menu not found", "not_found", http.StatusNotFound}
	ErrMenuCreateFailed = &AppError{"MENU_CREATE_FAILED", "Failed to create menu", "internal", http.StatusInternalServerError}
	ErrMenuUpdateFailed = &AppError{"MENU_UPDATE_FAILED", "Failed to update menu item", "internal", http.StatusInternalServerError}
	ErrMenuDeleteFailed = &AppError{"MENU_DELETE_FAILED", "Failed to delete menu item", "internal", http.StatusInternalServerError}
)

// merchant-related errors
var (
	ErrMerchantCreateFailed = &AppError{"MERCHANT_CREATE_FAILED", "Failed to create merchant", "internal", http.StatusInternalServerError}
	ErrMerchantNotFound     = &AppError{"MERCHANT_NOT_FOUND", "Merchant not found", "not_found", http.StatusNotFound}
	ErrNotMerchant          = &AppError{"NOT_MERCHANT", "User is not a merchant", "forbidden", http.StatusForbidden}
	ErrMerchantExists       = &AppError{"MERCHANT_EXISTS", "Merchant profile already exists", "conflict", http.StatusConflict}
	ErrMerchantUpdateFailed = &AppError{"MERCHANT_UPDATE_FAILED", "Failed to update merchant profile", "internal", http.StatusInternalServerError}
	ErrMerchantDeleteFailed = &AppError{"MERCHANT_DELETE_FAILED", "Failed to delete merchant ", "internal", http.StatusInternalServerError}
)

// authorization errors
var (
	ErrUnauthorized = &AppError{"UNAUTHORIZED", "Unauthorized access", "unauthorized", http.StatusUnauthorized}
	ErrForbidden    = &AppError{"FORBIDDEN", "Access forbidden", "forbidden", http.StatusForbidden}
	ErrInvalidToken = &AppError{"INVALID_TOKEN", "Invalid authentication token", "unauthorized", http.StatusUnauthorized}
)

// Validation errors
var (
	ErrValidationFailed = &AppError{"VALIDATION_FAILED", "Validation failed", "validation", http.StatusBadRequest}
	ErrInvalidInput     = &AppError{"INVALID_INPUT", "Invalid input provided", "validation", http.StatusBadRequest}
)

// Helper function to check if error is AppError
func IsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}

// Helper function to create custom validation error
func NewValidationError(message string) *AppError {
	return &AppError{
		Code:       "VALIDATION_ERROR",
		Message:    message,
		Type:       "validation",
		HTTPStatus: http.StatusBadRequest,
	}
}

// Helper function to create custom internal error
func NewInternalError(message string) *AppError {
	return &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    message,
		Type:       "internal",
		HTTPStatus: http.StatusInternalServerError,
	}
}
