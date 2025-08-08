package repository

import (
	db "task/db"
	"task/domain"
)

type EconomicIndicatorRepository struct {
	db *db.DB
}

func NewEconomicIndicatorRepository(db *db.DB) *EconomicIndicatorRepository {
	return &EconomicIndicatorRepository{db: db}
}
func (r *EconomicIndicatorRepository) GetEconomicIndicators() (domain.APIResponse, error) {
	return r.db.GetEconomicIndicators()
}
