package handlers

import (
	"net/http"
	"task/domain"
	"task/usecases"
	"task/utils"
)

type StockHandler struct {
	stockUseCase usecases.StockUseCase
}

func NewStockHandler(useCase usecases.StockUseCase) *StockHandler {
	return &StockHandler{
		stockUseCase: useCase,
	}
}

func (h *StockHandler) GetAllStocks(w http.ResponseWriter, r *http.Request) {
	filterType := r.URL.Query().Get("type")
	stock, err := h.stockUseCase.GetAllStocks()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	var stocks []domain.Stock

	switch filterType {
	case "gainer":
		stocks = filterGainers(stock)
	case "loser":
		stocks = filterLosers(stock)
	default:
		stocks = stock
	}

	utils.WriteJSON(w, http.StatusOK, stocks)
}
func filterGainers(stocks []domain.Stock) []domain.Stock {
	var gainers []domain.Stock
	for _, s := range stocks {
		if s.ChangePercent > 0 {
			gainers = append(gainers, s)
		}
	}
	return gainers
}

func filterLosers(stocks []domain.Stock) []domain.Stock {
	var losers []domain.Stock
	for _, s := range stocks {
		if s.ChangePercent < 0 {
			losers = append(losers, s)
		}
	}
	return losers
}
