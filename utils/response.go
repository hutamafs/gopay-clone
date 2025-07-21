package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SuccessResponse(c echo.Context, statusCode int, message string, data any) error {
	return c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string, err error) error {
	response := APIResponse{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	return c.JSON(statusCode, response)
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
