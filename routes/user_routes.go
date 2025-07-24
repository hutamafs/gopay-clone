package routes

import (
	"gopay-clone/config"
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(api *echo.Group, db *config.Database, jwtMiddleware echo.MiddlewareFunc) {
	userService := services.NewUserService(db)
	accountService := services.NewAccountService(db)
	orderService := services.NewOrderService(db)
	userHandler := handlers.NewUserHandler(userService, accountService, orderService)

	users := api.Group("/users")
	publicUsers := api.Group("/public/users")

	// PUBLIC routes (no middleware)
	publicUsers.POST("", userHandler.CreateUser)
	publicUsers.POST("/login", userHandler.Login)
	users.GET("", userHandler.GetAllUsers)

	// PROTECTED routes (with middleware)
	users.Use(jwtMiddleware)
	users.GET("/:id", userHandler.GetUserById)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
	users.GET("/:user_id/accounts", userHandler.GetAccountsByUser)
	users.GET("/:user_id/orders", userHandler.GetAllOrdersByUser)
}
