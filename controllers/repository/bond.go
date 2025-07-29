package repository

import (
	db "task/db"
	"task/domain"
	"time"
)

type BondRepository struct {
	db *db.DB
}

func NewBondRepository(database *db.DB) *BondRepository {
	return &BondRepository{
		db: database,
	}
}

func (r *BondRepository) GetBonds(search domain.BondSearchModel) ([]domain.Bond, error) {
	bonds := r.db.GetAllBonds()
	layout := "2006-01-02"
	var filteredBonds []domain.Bond
	for _, bond := range bonds {
		maturity, err := time.Parse(layout, bond.MaturityDate)
		if err != nil {
			continue
		}

		if search.MinCoupon != nil && bond.CouponRate < *search.MinCoupon {
			continue
		}
		if search.MaxCoupon != nil && bond.CouponRate > *search.MaxCoupon {
			continue
		}
		if search.MaturityBefore != nil && maturity.After(*search.MaturityBefore) {
			continue
		}
		if search.MaturityAfter != nil && maturity.Before(*search.MaturityAfter) {
			continue
		}

		filteredBonds = append(filteredBonds, bond)
	}

	return filteredBonds, nil
}
