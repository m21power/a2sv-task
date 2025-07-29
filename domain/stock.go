package domain

type Stock struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	ChangePercent float64 `json:"changePercent"`
	Volume        int     `json:"volume"`
}

type StockRepository interface {
	GetAllStocks() ([]Stock, error)
}
