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
    // Database configuration
	config := map[string]string{
		"host":     "localhost",  
		"port":     "5432",     
		"user":     "postgres",  
		"password": "postgres", 
		"dbname":   "workout",
		"sslmode":  "disable",
		"TimeZone": "UTC",
	}

	if os.Getenv("DB_HOST") != "" {
			config["host"] = os.Getenv("DB_HOST")
			config["port"] = os.Getenv("DB_PORT")
			config["user"] = os.Getenv("DB_USER")
			config["password"] = os.Getenv("DB_PASSWORD")
			config["dbname"] = os.Getenv("DB_NAME")
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
			panic("failed to connect database")
    }

    if err := sqlDB.Ping(); err != nil {
			log.Fatal("Failed to ping database:", err)
    }

    fmt.Println("Successfully connected to PostgreSQL 16!")

    return &Database{db}
}