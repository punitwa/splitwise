// store/expense_store.go

package store

import (
	"splitwise/models"

	"gorm.io/gorm"
)

type Expense interface {
	CreateExpense(expense models.Expense) (models.Expense, error)
	GetExpense(id int) (models.Expense, error)
}

type expense struct {
	db *gorm.DB
}

func NewExpense(db *gorm.DB) Expense {
	return &expense{db}
}

func (s *expense) CreateExpense(expense models.Expense) (models.Expense, error) {
	if err := s.db.Create(&expense).Error; err != nil {
		return models.Expense{}, err
	}
	return expense, nil
}

func (s *expense) GetExpense(id int) (models.Expense, error) {
	var expense models.Expense
	if err := s.db.Preload("Users").Preload("Settlements").First(&expense, id).Error; err != nil {
		return models.Expense{}, err
	}
	return expense, nil
}
