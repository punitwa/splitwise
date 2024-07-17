package services

import (
	"splitwise/models"
	"splitwise/store"
)

type ExpenseService interface {
	CreateExpense(expense models.Expense) (models.Expense, error)
	GetExpense(id int) (models.Expense, error)
}

type expenseService struct {
	store store.Store
}

func NewExpenseService(store store.Store) ExpenseService {
	return &expenseService{store}
}

func (s *expenseService) CreateExpense(expense models.Expense) (models.Expense, error) {
	return s.store.Expense().CreateExpense(expense)
}

func (s *expenseService) GetExpense(id int) (models.Expense, error) {
	return s.store.Expense().GetExpense(id)
}
