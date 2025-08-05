package repository

import (
	"task/db"
	"task/domain"
)

type PortRecomRepository struct {
	db *db.DB
}

func NewPortRecomRepository(database *db.DB) *PortRecomRepository {
	return &PortRecomRepository{
		db: database,
	}
}

func (r *PortRecomRepository) RecommendPortfolio(profile domain.PortfolioRequest) (domain.PortfolioResponse, error) {
	// Asset return rates
	const (
		stockRate = 0.20
		bondRate  = 0.10
		cashRate  = 0.07
	)

	var stockPerc, bondPerc, cashPerc float64

	// Allocation logic based on profile
	if profile.Profile == "active" {
		stockPerc = 0.6
		bondPerc = 0.3
		cashPerc = 0.1
	} else { // passive
		stockPerc = 0.3
		bondPerc = 0.5
		cashPerc = 0.2
	}

	// Investment amount calculation
	stockAmount := profile.InitialCapital * stockPerc
	bondAmount := profile.InitialCapital * bondPerc
	cashAmount := profile.InitialCapital * cashPerc

	// Yield calculations
	stockYield := stockAmount * stockRate
	bondYield := bondAmount * bondRate
	cashYield := cashAmount * cashRate

	totalReturn := stockYield + bondYield + cashYield
	finalAmount := profile.InitialCapital + totalReturn
	goalMet := finalAmount >= profile.TargetGoal
	gap := profile.TargetGoal - finalAmount

	// Response structure
	resp := domain.PortfolioResponse{
		Stock: domain.Asset{
			Name:   "EthioTelecom",
			Amount: stockAmount,
			Yield:  stockYield,
		},
		Bond: domain.Asset{
			Name:   "Ethiopian Gov Bond",
			Amount: bondAmount,
			Yield:  bondYield,
		},
		Cash: domain.Asset{
			Name:   "Birr",
			Amount: cashAmount,
			Yield:  cashYield,
		},
		TotalReturn: totalReturn,
		GoalMet:     goalMet,
		GapToGoal:   gap,
	}
	return resp, nil
}
