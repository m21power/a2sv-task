package usecases

import "task/domain"

type EconomicIndicatorUseCase struct {
	repo domain.EconomicIndicatorRepository
}

func NewEconomicIndicatorUseCase(repo domain.EconomicIndicatorRepository) *EconomicIndicatorUseCase {
	return &EconomicIndicatorUseCase{repo: repo}
}
func (uc *EconomicIndicatorUseCase) GetEconomicIndicators() (domain.APIResponse, error) {
	return uc.repo.GetEconomicIndicators()
}
