package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"splitwise/models"
	"splitwise/services"
)

type SettlementHandler struct {
	SettlementService services.SettlementService
}

func NewSettlementHandler(settlementService services.SettlementService) *SettlementHandler {
	return &SettlementHandler{SettlementService: settlementService}
}

func (h *SettlementHandler) CreateSettlement(w http.ResponseWriter, r *http.Request) {
	var settlement models.Settlement
	if err := json.NewDecoder(r.Body).Decode(&settlement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdSettlement, err := h.SettlementService.CreateSettlement(settlement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdSettlement)
}

func (h *SettlementHandler) GetSettlement(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid settlement ID", http.StatusBadRequest)
		return
	}
	settlement, err := h.SettlementService.GetSettlement(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settlement)
}

// Add additional methods for PUT and DELETE if necessary
