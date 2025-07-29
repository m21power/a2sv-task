package usecases

import "task/domain"

type BondUseCase struct {
	bondRepo domain.BondRepository
}

func NewBondUseCase(repo domain.BondRepository) BondUseCase {
	return BondUseCase{
		bondRepo: repo,
	}
}
func (u *BondUseCase) GetBonds(search domain.BondSearchModel) ([]domain.Bond, error) {
	return u.bondRepo.GetBonds(search)
}
