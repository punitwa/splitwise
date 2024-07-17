package store

import "gorm.io/gorm"

type Store interface {
	User() User
	Expense() Expense
	Group() Group
	Settlement() Settlement
}

type store struct {
	db         *gorm.DB
	user       User
	expense    Expense
	group      Group
	settlement Settlement
}

func NewStore(db *gorm.DB) Store {
	return &store{
		db:         db,
		user:       NewUser(db),
		expense:    NewExpense(db),
		group:      NewGroup(db),
		settlement: NewSettlement(db),
	}
}

func (s *store) User() User {
	return s.user
}

func (s *store) Expense() Expense {
	return s.expense
}

func (s *store) Group() Group {
	return s.group
}

func (s *store) Settlement() Settlement {
	return s.settlement
}
