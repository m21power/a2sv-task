package domain

import "time"

type Bond struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Issuer       string  `json:"issuer"`
	CouponRate   float64 `json:"couponRate"`
	MaturityDate string  `json:"maturityDate"`
	Price        float64 `json:"price"`
}

type BondSearchModel struct {
	MinCoupon      *float64   `json:"minCoupon,omitempty"`
	MaxCoupon      *float64   `json:"maxCoupon,omitempty"`
	MaturityBefore *time.Time `json:"maturityBefore,omitempty"`
	MaturityAfter  *time.Time `json:"maturityAfter,omitempty"`
}

type BondRepository interface {
	GetBonds(search BondSearchModel) ([]Bond, error)
}
