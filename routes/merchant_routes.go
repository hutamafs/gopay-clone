package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterMerchantRoutes(api *echo.Group, db *config.Database, jwtMiddleware echo.MiddlewareFunc) {
	merchantService := services.NewMerchantService(db)
	userService := services.NewUserService(db)
	merchantHandler := handlers.NewMerchantHandler(userService, merchantService)

	publicMerchantAPI := api.Group("/merchants")
	merchants := api.Group("/merchants")
	merchants.Use(jwtMiddleware)
	{
		publicMerchantAPI.POST("", merchantHandler.CreateMerchant)
		merchants.GET("", merchantHandler.GetAllMerchants, jwtMiddleware)
		merchants.GET("/:merchant_id", merchantHandler.GetAllMerchants, jwtMiddleware)
		merchants.PUT("/:merchant_id", merchantHandler.UpdateMerchantByID, jwtMiddleware)
	}
}
