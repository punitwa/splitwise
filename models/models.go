package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type Expense struct {
	gorm.Model
	Description string
	Amount      float64
	PayerID     int
	SplitType   string
	Users       []User `gorm:"many2many:expense_users;"`
	Settlements []Settlement
}

type Settlement struct {
	gorm.Model
	ExpenseID int
	Expense   Expense
	PayerID   int
	PayeeID   int
	Amount    float64
	Status    string
}

type Group struct {
	gorm.Model
	Name     string
	Members  []User `gorm:"many2many:group_users;"`
	Expenses []Expense
}
