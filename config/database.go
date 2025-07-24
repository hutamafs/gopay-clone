package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func InitDatabase() *Database {
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		fmt.Println("Using DATABASE_URL for connection")
		db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatal("Failed to connect to database with DATABASE_URL:", err)
		}

		// Test the connection
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Failed to get database instance:", err)
		}

		if err := sqlDB.Ping(); err != nil {
			log.Fatal("Failed to ping database:", err)
		}

		fmt.Println("Successfully connected to PostgreSQL via DATABASE_URL!")
		return &Database{db}
	}

	// Database configuration
	config := map[string]string{
		"host":     getEnvOrDefault("DB_HOST", "localhost"),
		"port":     getEnvOrDefault("DB_PORT", "5432"),
		"user":     getEnvOrDefault("DB_USER", "postgres"),
		"password": getEnvOrDefault("DB_PASSWORD", "postgres"),
		"dbname":   getEnvOrDefault("DB_NAME", "gopay"),
		"sslmode":  getEnvOrDefault("DB_SSLMODE", "disable"),
		"TimeZone": "UTC",
	}
	// Create connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config["host"], config["user"], config["password"],
		config["dbname"], config["port"], config["sslmode"], config["TimeZone"])

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL logging
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL 16!")

	return &Database{db}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
