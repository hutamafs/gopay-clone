package routes

import (
	"gopay-clone/handlers"
	"gopay-clone/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func RegisterUserRoutes(api *echo.Group, db *gorm.DB) {
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	users := api.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUserById)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}