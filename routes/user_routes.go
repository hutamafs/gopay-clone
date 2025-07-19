package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo"
)

func RegisterUserRoutes(api *echo.Group, db *config.Database) {
	userService := services.NewUserService(db)
	accountService := services.NewAccountService(db)
	userHandler := handlers.NewUserHandler(userService, accountService)

	users := api.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUserById)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)

		users.GET("/:user_id/accounts", userHandler.GetAccountsByUser)
	}
}
