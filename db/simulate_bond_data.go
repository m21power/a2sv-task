package db

import "task/domain"

var simulateBond = []domain.Bond{
	{
		ID:           "BND001",
		Name:         "Government Bond A",
		Issuer:       "Ministry of Finance",
		CouponRate:   5.25,
		MaturityDate: "2027-06-30",
		Price:        102.15,
	},
	{
		ID:           "BND002",
		Name:         "Infrastructure Bond X",
		Issuer:       "Ethiopian Roads Corp",
		CouponRate:   6.00,
		MaturityDate: "2030-12-31",
		Price:        101.80,
	},
	{
		ID:           "BND003",
		Name:         "Health Sector Bond",
		Issuer:       "Health Ministry",
		CouponRate:   3.75,
		MaturityDate: "2025-05-15",
		Price:        98.40,
	},
	{
		ID:           "BND004",
		Name:         "Green Energy Bond",
		Issuer:       "Ethiopian Energy Corp",
		CouponRate:   4.75,
		MaturityDate: "2026-12-15",
		Price:        99.50,
	},
	{
		ID:           "BND005",
		Name:         "Agricultural Bond",
		Issuer:       "Ministry of Agriculture",
		CouponRate:   5.00,
		MaturityDate: "2029-03-01",
		Price:        100.00,
	},
	{
		ID:           "BND006",
		Name:         "Diaspora Bond",
		Issuer:       "National Bank of Ethiopia",
		CouponRate:   6.50,
		MaturityDate: "2032-09-01",
		Price:        105.25,
	},
	{
		ID:           "BND007",
		Name:         "Railway Project Bond",
		Issuer:       "Transport Authority",
		CouponRate:   4.00,
		MaturityDate: "2028-06-01",
		Price:        97.75,
	},
	{
		ID:           "BND008",
		Name:         "Tech Infrastructure Bond",
		Issuer:       "ICT Development Fund",
		CouponRate:   4.90,
		MaturityDate: "2027-11-30",
		Price:        99.20,
	},
	{
		ID:           "BND009",
		Name:         "Private Sector Bond",
		Issuer:       "Ethiopian Investment Bank",
		CouponRate:   5.30,
		MaturityDate: "2026-08-10",
		Price:        100.10,
	},
	{
		ID:           "BND010",
		Name:         "Emergency Relief Bond",
		Issuer:       "Disaster Management Agency",
		CouponRate:   3.50,
		MaturityDate: "2025-02-20",
		Price:        96.90,
	}}

func (db *DB) GetAllBonds() []domain.Bond {
	return simulateBond
}
