package main

import (
	"fmt"
	"gopay-clone/config"
	"gopay-clone/middleware"
	"gopay-clone/models"
	"gopay-clone/routes"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func setupRoutes(e *echo.Echo, db *gorm.DB) {
	api := e.Group("api/v1")

	// register route groups
	routes.RegisterUserRoutes(api, db)
}

func main() {
	// database
	db := config.InitDatabase()

	// migrate user
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.QrCode{})
	db.AutoMigrate(&models.Transaction{})

	// echo
	e := echo.New()

	// apply middleware
	middleware.ApplyMiddleware(e)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":   "healthy",
			"database": "connected",
		})
	})

	setupRoutes(e, db.DB)

	// Start server
	fmt.Println("Server starting on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
