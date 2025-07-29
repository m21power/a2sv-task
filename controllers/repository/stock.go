package repository

import (
	db "task/db"
	"task/domain"
)

type StockRepository struct {
	db *db.DB
}

func NewStockRepository(database *db.DB) *StockRepository {
	return &StockRepository{
		db: database,
	}
}
func (r *StockRepository) GetAllStocks() ([]domain.Stock, error) {
	return r.db.GetAllStocks(), nil
}
