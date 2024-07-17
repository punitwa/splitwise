package store

import (
	"splitwise/models"

	"gorm.io/gorm"
)

type Settlement interface {
	CreateSettlement(settlement models.Settlement) (models.Settlement, error)
	GetSettlement(id int) (models.Settlement, error)
}

type settlement struct {
	db *gorm.DB
}

func NewSettlement(db *gorm.DB) Settlement {
	return &settlement{db}
}

func (s *settlement) CreateSettlement(settlement models.Settlement) (models.Settlement, error) {
	if err := s.db.Create(&settlement).Error; err != nil {
		return models.Settlement{}, err
	}
	return settlement, nil
}

func (s *settlement) GetSettlement(id int) (models.Settlement, error) {
	var settlement models.Settlement
	if err := s.db.First(&settlement, id).Error; err != nil {
		return models.Settlement{}, err
	}
	return settlement, nil
}
