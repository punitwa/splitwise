package main

import (
	"log"
	"net/http"
	"splitwise/api/routes"
	"splitwise/config"
	"splitwise/services"
	"splitwise/store"
)

func main() {
	// Load configuration
	if err := config.LoadConfig("config.json"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to the database
	config.ConnectDatabase()
	db := config.DB

	// Run migrations
	// migrations.Migrate(db)

	// Seed data (if needed)
	// migrations.Seed(db)

	// Initialize store
	store := store.NewStore(db)

	// Initialize services
	userService := services.NewUserService(store)
	expenseService := services.NewExpenseService(store)
	groupService := services.NewGroupService(store)
	settlementService := services.NewSettlementService(store)

	// Setup routes
	router := routes.SetupRoutes(userService, expenseService, groupService, settlementService)

	// Start server
	serverConfig := config.AppConfig.Server
	log.Printf("Server running at :%s", serverConfig.Port)
	http.ListenAndServe(":"+serverConfig.Port, router)
}
