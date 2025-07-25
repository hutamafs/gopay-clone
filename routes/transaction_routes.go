package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterTransactionRoutes(api *echo.Group, db *config.Database, jwtMiddleware echo.MiddlewareFunc) {
	transactionService := services.NewTransactionService(db)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	transactions := api.Group("/transactions")
	transactions.Use(jwtMiddleware)
	{
		transactions.POST("", transactionHandler.CreateTransaction)
		transactions.GET("/:transaction_id", transactionHandler.GetTransactionDetail)
		transactions.PUT("/:transaction_id", transactionHandler.UpdateTransactionDetail)
	}
}
