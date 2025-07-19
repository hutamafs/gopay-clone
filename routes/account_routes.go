package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo"
)

func RegisterAccountRoutes(api *echo.Group, db *config.Database) {
	accountService := services.NewAccountService(db)
	accountHandler := handlers.NewAccountHandler(accountService)

	accounts := api.Group("/accounts")
	{
		accounts.POST("", accountHandler.CreateAccount)
		accounts.GET("/:account_id/balance", accountHandler.GetBalanceByAccountId)
		accounts.PUT("/:account_id", accountHandler.UpdateAccount)
		accounts.GET("/:account_id/detail", accountHandler.GetAccountDetail)
	}
}
