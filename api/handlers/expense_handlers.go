// api/handlers/expense_handler.go

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"splitwise/models"
	"splitwise/services"
)

type ExpenseHandler struct {
	ExpenseService services.ExpenseService
}

func NewExpenseHandler(expenseService services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{ExpenseService: expenseService}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdExpense, err := h.ExpenseService.CreateExpense(expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdExpense)
}

func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid expense ID", http.StatusBadRequest)
		return
	}
	expense, err := h.ExpenseService.GetExpense(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}

// Add additional methods for PUT and DELETE if necessary
