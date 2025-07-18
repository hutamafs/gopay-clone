package utils

import (
	"github.com/labstack/echo"
)

func BindAndValidate[T any](c echo.Context, req *T, validatorFn func(*T) error) error {
	if err := c.Bind(req); err != nil {
		ValidationErrorResponse(c, err)
		return err
	}

	if err := validatorFn(req); err != nil {
		ValidationErrorResponse(c, err)
		return err
	}
	return nil
}