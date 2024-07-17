// store/group_store.go

package store

import (
	"splitwise/models"

	"gorm.io/gorm"
)

type Group interface {
	CreateGroup(group models.Group) (models.Group, error)
	GetGroup(id int) (models.Group, error)
	AddExpenseToGroup(groupID int, expense models.Expense) (models.Group, error)
}

type group struct {
	db *gorm.DB
}

func NewGroup(db *gorm.DB) Group {
	return &group{db}
}

func (s *group) CreateGroup(group models.Group) (models.Group, error) {
	if err := s.db.Create(&group).Error; err != nil {
		return models.Group{}, err
	}
	return group, nil
}

func (s *group) GetGroup(id int) (models.Group, error) {
	var group models.Group
	if err := s.db.Preload("Members").Preload("Expenses").First(&group, id).Error; err != nil {
		return models.Group{}, err
	}
	return group, nil
}

func (s *group) AddExpenseToGroup(groupID int, expense models.Expense) (models.Group, error) {
	var group models.Group
	if err := s.db.First(&group, groupID).Error; err != nil {
		return models.Group{}, err
	}

	if err := s.db.Create(&expense).Error; err != nil {
		return models.Group{}, err
	}

	group.Expenses = append(group.Expenses, expense)
	if err := s.db.Save(&group).Error; err != nil {
		return models.Group{}, err
	}
	return group, nil
}
