package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterOrderRoutes(api *echo.Group, db *config.Database, jwtMiddleware echo.MiddlewareFunc) {
	orderService := services.NewOrderService(db)
	merchantService := services.NewMerchantService(db)
	userService := services.NewUserService(db)

	orderHandler := handlers.NewOrderHandler(orderService, merchantService, userService)

	orders := api.Group("/orders")
	orders.Use(jwtMiddleware)
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("/:order_id", orderHandler.GetOrderByID)
	}
}
