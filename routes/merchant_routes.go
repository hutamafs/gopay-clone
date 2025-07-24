package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterMerchantRoutes(api *echo.Group, db *config.Database, jwtMiddleware echo.MiddlewareFunc) {
	merchantService := services.NewMerchantService(db)
	menuService := services.NewMenuItemService(db)
	userService := services.NewUserService(db)
	merchantHandler := handlers.NewMerchantHandler(userService, merchantService)
	menuHandler := handlers.NewMenuHandler(menuService, merchantService)

	publicMerchantAPI := api.Group("/public/merchants")
	merchants := api.Group("/merchants")
	menus := api.Group("/menus")
	merchants.Use(jwtMiddleware)
	menus.Use(jwtMiddleware)
	{
		publicMerchantAPI.POST("", merchantHandler.CreateMerchant)
		merchants.GET("", merchantHandler.GetAllMerchants)
		merchants.GET("/:merchant_id", merchantHandler.GetAllMerchants)
		merchants.PUT("/:merchant_id", merchantHandler.UpdateMerchantByID)

		// menu item
		merchants.POST("/:merchant_id/menu-item", menuHandler.CreateMenu)
		merchants.GET("/:merchant_id/menu-item", menuHandler.GetAllMenus)
		merchants.GET("/menu-item/:menu_id", menuHandler.GetMenuByID)
		merchants.PUT("/:merchant_id/menu-item/:menu_id", menuHandler.UpdateMenuItem)
		merchants.DELETE("/:merchant_id/menu-item/:menu_id", menuHandler.DeleteMenuItem)

		// get all menus by filter
		menus.GET("/menu-items", menuHandler.GetAllMenus)
	}
}
