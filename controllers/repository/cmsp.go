package repository

import (
	"errors"
	"fmt"
	db "task/db"
	"task/domain"
	"time"
)

type CMSPRepository struct {
	db *db.DB
}

func NewCMSPRepository(db *db.DB) *CMSPRepository {
	return &CMSPRepository{db: db}
}

func (r *CMSPRepository) GetAllCMSPs(filter domain.FilterModel) ([]domain.CMSP, error) {
	if filter.Type == "" && filter.LicensedBefore == nil && filter.LicensedAfter == nil {
		return r.db.GetAllCMSPs(), nil
	}
	return r.GetCMSPByFilter(filter)
}
func (r *CMSPRepository) GetCMSPByID(id int) (*domain.CMSP, error) {
	cmspList := r.db.GetAllCMSPs()
	for _, cmsp := range cmspList {
		if cmsp.ID == id {
			return &cmsp, nil
		}
	}
	return nil, errors.New("CMSP not found")
}

func (r *CMSPRepository) GetCMSPByFilter(filter domain.FilterModel) ([]domain.CMSP, error) {
	var filteredCMSPs []domain.CMSP
	fmt.Println("Filtering CMSPs with filter:", filter.Type, filter.LicensedBefore, filter.LicensedAfter)
	cmspList := r.db.GetAllCMSPs()
	for _, cmsp := range cmspList {
		licensedDate, err := time.Parse("2006-01-02", cmsp.LicensedDate)
		if err != nil {
			continue
		}
		if (filter.Type == "" || cmsp.Type == filter.Type) &&
			(filter.LicensedBefore == nil || licensedDate.Before(*filter.LicensedBefore)) &&
			(filter.LicensedAfter == nil || licensedDate.After(*filter.LicensedAfter)) {
			filteredCMSPs = append(filteredCMSPs, cmsp)
		}
	}
	return filteredCMSPs, nil
}
