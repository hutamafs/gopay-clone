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
	userHandler := handlers.NewUserHandler(userService, accountService)

	users := api.Group("/users")
	publicUsers := api.Group("/public/users")

	// PUBLIC routes (no middleware)
	publicUsers.POST("", userHandler.CreateUser)
	publicUsers.POST("/login", userHandler.Login)
	users.GET("", userHandler.GetAllUsers)

	// PROTECTED routes (with middleware)
	users.Use(jwtMiddleware)
	users.GET("/:id", userHandler.GetUserById, jwtMiddleware)
	users.PUT("/:id", userHandler.UpdateUser, jwtMiddleware)
	users.DELETE("/:id", userHandler.DeleteUser, jwtMiddleware)
	users.GET("/:user_id/accounts", userHandler.GetAccountsByUser, jwtMiddleware)
}
