package migrations

import (
	"log"
	"splitwise/models"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("Seeding database...")

	// Seed users
	users := []models.User{
		{Name: "Alice", Email: "alice@example.com", Password: "password"},
		{Name: "Bob", Email: "bob@example.com", Password: "password"},
	}
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed users: %v", err)
		}
	}

	// Seed groups
	groups := []models.Group{
		{Name: "Trip to Paris"},
	}
	for _, group := range groups {
		if err := db.Create(&group).Error; err != nil {
			log.Fatalf("Failed to seed groups: %v", err)
		}
	}

	// Seed expenses
	expenses := []models.Expense{
		{Description: "Flight Tickets", Amount: 500.0, PayerID: 1, SplitType: "Equal"},
	}
	for _, expense := range expenses {
		if err := db.Create(&expense).Error; err != nil {
			log.Fatalf("Failed to seed expenses: %v", err)
		}
	}

	// Seed settlements
	settlements := []models.Settlement{
		{ExpenseID: 1, PayerID: 1, PayeeID: 2, Amount: 250.0, Status: "Pending"},
	}
	for _, settlement := range settlements {
		if err := db.Create(&settlement).Error; err != nil {
			log.Fatalf("Failed to seed settlements: %v", err)
		}
	}

	log.Println("Database seeded.")
}
