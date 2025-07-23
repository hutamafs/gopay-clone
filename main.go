package main

import (
	"gopay-clone/config"
	"gopay-clone/migrations"
	"gopay-clone/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo, db *config.Database, secret string) {
	api := e.Group("/api/v1")

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(secret),
		ContextKey:  "user", // This stores the *jwt.Token in context
		TokenLookup: "header:Authorization:Bearer ",
	})
	// Register routes
	routes.RegisterUserRoutes(api, db, jwtMiddleware)
	routes.RegisterMerchantRoutes(api, db, jwtMiddleware)
	routes.RegisterAccountRoutes(api, db, jwtMiddleware)
	routes.RegisterTransactionRoutes(api, db, jwtMiddleware)
	routes.RegisterQRRoutes(api, db, jwtMiddleware)
	routes.RegisterOrderRoutes(api, db, jwtMiddleware)
	routes.RegisterDriverRoutes(api, db, jwtMiddleware)
}

func main() {
	_ = godotenv.Load()
	secret := os.Getenv("JWT_SECRET")

	// database
	db := config.InitDatabase()

	// migrate models
	migrations.RunMigration(db)

	// echo
	e := echo.New()

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":   "healthy",
			"database": "connected",
		})
	})

	setupRoutes(e, db, secret)

	e.Logger.Fatal(e.Start(":8080"))
}
