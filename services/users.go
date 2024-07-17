package services

import (
	"splitwise/models"
	"splitwise/store"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id int) (models.User, error)
}

type userService struct {
	store store.Store
}

func NewUserService(store store.Store) UserService {
	return &userService{store}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.store.User().CreateUser(user)
}

func (s *userService) GetUser(id int) (models.User, error) {
	return s.store.User().GetUser(id)
}
