package main

import (
	"gopay-clone/config"
	_ "gopay-clone/docs"
	"gopay-clone/migrations"
	"gopay-clone/routes"
	"net/http"
	"os"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo, db *config.Database, secret string) {
	api := e.Group("/api/v1")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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

// @title GoClone API
// @version 1.0
// @description GoClone super app API
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
