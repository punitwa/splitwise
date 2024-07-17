package routes

import (
	"splitwise/api/handlers"
	"splitwise/services"

	"github.com/gorilla/mux"
)

func SetupRoutes(userService services.UserService, expenseService services.ExpenseService, groupService services.GroupService, settlementService services.SettlementService) *mux.Router {
	router := mux.NewRouter()

	userHandler := handlers.NewUserHandler(userService)
	expenseHandler := handlers.NewExpenseHandler(expenseService)
	groupHandler := handlers.NewGroupHandler(groupService)
	settlementHandler := handlers.NewSettlementHandler(settlementService)

	api := router.PathPrefix("/api").Subrouter()
	// api.Use(middleware.Auth)

	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	// Add additional routes for PUT and DELETE

	api.HandleFunc("/expenses", expenseHandler.CreateExpense).Methods("POST")
	api.HandleFunc("/expenses", expenseHandler.GetExpense).Methods("GET")
	// Add additional routes for PUT and DELETE

	api.HandleFunc("/groups", groupHandler.CreateGroup).Methods("POST")
	api.HandleFunc("/groups", groupHandler.GetGroup).Methods("GET")
	// Add additional routes for PUT and DELETE

	api.HandleFunc("/settlements", settlementHandler.CreateSettlement).Methods("POST")
	api.HandleFunc("/settlements", settlementHandler.GetSettlement).Methods("GET")
	// Add additional routes for PUT and DELETE

	return router
}
