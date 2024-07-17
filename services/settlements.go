package services

import (
	"splitwise/models"
	"splitwise/store"
)

type SettlementService interface {
	CreateSettlement(settlement models.Settlement) (models.Settlement, error)
	GetSettlement(id int) (models.Settlement, error)
}

type settlementService struct {
	store store.Store
}

func NewSettlementService(store store.Store) SettlementService {
	return &settlementService{store}
}

func (s *settlementService) CreateSettlement(settlement models.Settlement) (models.Settlement, error) {
	return s.store.Settlement().CreateSettlement(settlement)
}

func (s *settlementService) GetSettlement(id int) (models.Settlement, error) {
	return s.store.Settlement().GetSettlement(id)
}
