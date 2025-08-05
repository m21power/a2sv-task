package domain

type PortfolioRequest struct {
	Profile        string  `json:"profile"`
	InitialCapital float64 `json:"initialCapital"`
	TargetGoal     float64 `json:"targetGoal"`
}

type Asset struct {
	Name   string  `json:"Name"`
	Amount float64 `json:"Amount"`
	Yield  float64 `json:"Yield"`
}

type PortfolioResponse struct {
	Stock       Asset   `json:"Stock"`
	Bond        Asset   `json:"Bond"`
	Cash        Asset   `json:"Cash"`
	TotalReturn float64 `json:"TotalReturn"`
	GoalMet     bool    `json:"GoalMet"`
	GapToGoal   float64 `json:"GapToGoal"`
}
type PortfolioRecommendationRepository interface {
	RecommendPortfolio(profile PortfolioRequest) (PortfolioResponse, error)
}
