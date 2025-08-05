package usecases

import "task/domain"

type CMSPUseCase struct {
	cmspRepository domain.CMSPRepository
}

func NewCMSPUseCase(repo domain.CMSPRepository) *CMSPUseCase {
	return &CMSPUseCase{cmspRepository: repo}
}

func (u *CMSPUseCase) GetAllCMSPs(filter domain.FilterModel) ([]domain.CMSP, error) {
	return u.cmspRepository.GetAllCMSPs(filter)
}
func (u *CMSPUseCase) GetCMSPByID(id int) (*domain.CMSP, error) {
	return u.cmspRepository.GetCMSPByID(id)
}
