package domain

type RawYearData struct {
	Year            int     `json:"year"`
	DepositRate     float64 `json:"deposit_rate"`
	InflationRate   float64 `json:"inflation_rate"`
	GDPGrowth       float64 `json:"gdp_growth"`
	ExchangeRate    float64 `json:"exchange_rate"`
	ForeignReserves float64 `json:"foreign_reserves"`
	GovtDebtGDP     float64 `json:"govt_debt_gdp"`
}
type IndicatorMeta struct {
	Name        string `json:"name"`
	Indicator   string `json:"indicator"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
}

type DerivedYearData struct {
	Year             int     `json:"year"`
	RealInterestRate float64 `json:"real_interest_rate"`
	GDPGrowth        float64 `json:"gdp_growth"`
	ExchangeRate     float64 `json:"exchange_rate"`
	ForeignReserves  float64 `json:"foreign_reserves"`
	GovtDebtGDP      float64 `json:"govt_debt_gdp"`
}

type APIResponse struct {
	Metadata []IndicatorMeta   `json:"metadata"`
	Data     []DerivedYearData `json:"data"`
	Source   string            `json:"source"`
}

type EconomicIndicatorRepository interface {
	GetEconomicIndicators() (APIResponse, error)
}
