package migrations

import (
	"log"
	"splitwise/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Expense{},
		&models.Settlement{},
		&models.Group{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migrations completed.")
}
