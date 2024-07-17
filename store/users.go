package store

import (
	"splitwise/models"

	"gorm.io/gorm"
)

type User interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id int) (models.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &user{db}
}

func (s *user) CreateUser(user models.User) (models.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *user) GetUser(id int) (models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
