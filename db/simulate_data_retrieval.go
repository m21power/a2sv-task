package db

import "task/domain"

var SimulateData = []domain.Stock{
	{
		Symbol:        "ETHB",
		Name:          "EthioBank",
		Price:         125.75,
		ChangePercent: 2.5,
		Volume:        150000,
	},
	{
		Symbol:        "ABCA",
		Name:          "Abyssinia Capital",
		Price:         87.30,
		ChangePercent: -1.2,
		Volume:        85000,
	},
	{
		Symbol:        "HBZ",
		Name:          "Habesha Breweries",
		Price:         102.40,
		ChangePercent: 0.8,
		Volume:        92000,
	},
	{
		Symbol:        "TMTE",
		Name:          "Telecom Ethiopia",
		Price:         150.10,
		ChangePercent: -0.6,
		Volume:        123000,
	},
	{
		Symbol:        "ADIS",
		Name:          "Addis Tech",
		Price:         200.00,
		ChangePercent: 4.1,
		Volume:        76000,
	},
}

type DB struct {
}

func NewRepository() *DB {
	return &DB{}
}

func (db *DB) GetAllStocks() []domain.Stock {
	return SimulateData
}
