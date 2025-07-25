package utils

import (
	apperrors "gopay-clone/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    any          `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
	Type    string `json:"type,omitempty"`
}

func SuccessResponse(c echo.Context, statusCode int, message string, data any) error {
	return c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// enhanced error response that handles AppError
func ErrorResponse(c echo.Context, statusCode int, message string, err error) error {
	response := APIResponse{
		Success: false,
		Message: message,
	}

	if err != nil {
		if appErr, ok := apperrors.IsAppError(err); ok {
			response.Error = &ErrorDetail{
				Code:    appErr.Code,
				Message: appErr.Message,
				Type:    appErr.Type,
			}
		} else {
			response.Error = &ErrorDetail{
				Message: err.Error(),
			}
		}
	}

	return c.JSON(statusCode, response)
}

// split error response that automatically determines status code from AppError
func SplitErrorResponse(c echo.Context, err error) error {
	if appErr, ok := apperrors.IsAppError(err); ok {
		return ErrorResponse(c, appErr.HTTPStatus, appErr.Message, appErr)
	}
	return InternalErrorResponse(c, err)
}

func ValidationErrorResponse(c echo.Context, err error) error {
	return ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
}

func NotFoundResponse(c echo.Context, resource string) error {
	return ErrorResponse(c, http.StatusNotFound, resource+" not found", nil)
}

func InternalErrorResponse(c echo.Context, err error) error {
	return ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err)
}

func ForbiddenResponse(c echo.Context, err error) error {
	return ErrorResponse(c, http.StatusForbidden, "Forbidden", err)
}

func UnauthorizedResponse(c echo.Context, err error) error {
	return ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", err)
}

func ConflictResponse(c echo.Context, message string, err error) error {
	return ErrorResponse(c, http.StatusConflict, message, err)
}
