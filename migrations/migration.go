package migrations

import (
	"fmt"
	"gopay-clone/config"
	"gopay-clone/models"
)

func RunMigration(db *config.Database) error {
	models := []interface{}{
		&models.User{},
		&models.Account{},
		&models.Contact{},
		&models.QrCode{},
		&models.Transaction{},
		&models.DriverProfile{},
		&models.MerchantProfile{},
		&models.MenuItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Ride{},
	}
	fmt.Println("Running database migrations...")

	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	fmt.Println("migration completed")
	return nil
}
