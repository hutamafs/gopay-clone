package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ApplyMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
			c.Response().Header().Set("X-API-Version", "v1")
			return next(c)
		}
	})
}