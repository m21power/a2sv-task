package usecases

import "task/domain"

type PortfolioRecommendationUseCase struct {
	portfolioRepo domain.PortfolioRecommendationRepository
}

func NewPortfolioRecommendationUseCase(repo domain.PortfolioRecommendationRepository) *PortfolioRecommendationUseCase {
	return &PortfolioRecommendationUseCase{portfolioRepo: repo}
}
func (u *PortfolioRecommendationUseCase) RecommendPortfolio(profile domain.PortfolioRequest) (domain.PortfolioResponse, error) {
	return u.portfolioRepo.RecommendPortfolio(profile)
}
