package usecases

import "task/domain"

type StockUseCase struct {
	stockRepo domain.StockRepository
}

func NewStockUseCase(repo domain.StockRepository) StockUseCase {
	return StockUseCase{
		stockRepo: repo,
	}
}

func (uc *StockUseCase) GetAllStocks() ([]domain.Stock, error) {
	return uc.stockRepo.GetAllStocks()
}
