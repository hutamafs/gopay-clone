package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterAccountRoutes(api *echo.Group, db *config.Database, jwtMiddleware echo.MiddlewareFunc) {
	accountService := services.NewAccountService(db)
	transactionService := services.NewTransactionService(db)
	accountHandler := handlers.NewAccountHandler(accountService, transactionService)

	accounts := api.Group("/accounts")
	accounts.Use(jwtMiddleware)
	{
		accounts.POST("", accountHandler.CreateAccount)
		accounts.GET("/:account_id/balance", accountHandler.GetBalanceByAccountId)
		accounts.PUT("/:account_id", accountHandler.UpdateAccount)
		accounts.GET("/:account_id/detail", accountHandler.GetAccountDetail)
		accounts.GET("/:account_id/transactions", accountHandler.GetTransactionByAccounts)
	}
}
