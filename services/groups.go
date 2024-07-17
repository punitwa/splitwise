package services

import (
	"splitwise/models"
	"splitwise/store"
)

type GroupService interface {
	CreateGroup(group models.Group) (models.Group, error)
	GetGroup(id int) (models.Group, error)
	AddExpenseToGroup(groupID int, expense models.Expense) (models.Group, error)
	AddMemberToGroup(groupID int, memberID int) (models.Group, error)
}

type groupService struct {
	store store.Store
}

func NewGroupService(store store.Store) GroupService {
	return &groupService{store}
}

func (s *groupService) CreateGroup(group models.Group) (models.Group, error) {
	return s.store.Group().CreateGroup(group)
}

func (s *groupService) GetGroup(id int) (models.Group, error) {
	return s.store.Group().GetGroup(id)
}

func (s *groupService) AddExpenseToGroup(groupID int, expense models.Expense) (models.Group, error) {
	return s.store.Group().AddExpenseToGroup(groupID, expense)
}

func (s *groupService) AddMemberToGroup(groupID int, memberID int) (models.Group, error) {
	return s.store.Group().AddMemberToGroup(groupID, memberID)
}
