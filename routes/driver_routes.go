package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterDriverRoutes(api *echo.Group, db *config.Database, jwtMiddleware echo.MiddlewareFunc) {
	driverService := services.NewDriverService(db)
	userService := services.NewUserService(db)
	driverHandler := handlers.NewDriverHandler(driverService, userService)

	drivers := api.Group("/drivers")

	publicDrivers := api.Group("/public/drivers")
	publicDrivers.GET("", driverHandler.GetAllDrivers)
	publicDrivers.POST("", driverHandler.CreateDriver)

	drivers.Use(jwtMiddleware)
	{
		drivers.GET("/available", driverHandler.GetAvailableDrivers)

		// driver profile management (for drivers themselves)
		drivers.GET("/:driver_id", driverHandler.GetDriverByID)
		drivers.PUT("/profile", driverHandler.UpdateDriverProfile)
		drivers.DELETE("/profile", driverHandler.DeleteDriverProfile)

		// driver status and location (for drivers themselves)
		drivers.PUT("/status", driverHandler.UpdateDriverStatus)
		drivers.PUT("/location", driverHandler.UpdateDriverLocation)
	}
}
