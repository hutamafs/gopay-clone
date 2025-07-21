package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterQRRoutes(api *echo.Group, db *config.Database) {
	qrService := services.NewQRService(db)
	transactionHandler := handlers.NewQRHandler(qrService)

	transactions := api.Group("/qr")
	{
		transactions.POST("", transactionHandler.CreateQR)
		transactions.PUT("/:qr_id", transactionHandler.ScanQr)
	}
}
