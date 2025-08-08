package db

import "task/domain"

var metadata = []domain.IndicatorMeta{
	{Name: "Real Interest Rate", Unit: "%", Description: "The nominal deposit rate minus the inflation rate; shows the real return on savings."},
	{Name: "GDP Growth", Unit: "%", Description: "Annual percentage growth rate of GDP at market prices based on constant local currency."},
	{Name: "Exchange Rate", Unit: "ETB/USD", Description: "Annual average official exchange rate of the Ethiopian Birr to the US Dollar."},
	{Name: "Foreign Reserves", Unit: "Months of import cover", Description: "Official foreign currency reserves expressed in months of import coverage."},
	{Name: "Government Debt to GDP", Unit: "%", Description: "Gross government debt as a percentage of GDP."},
}
var rawData = []domain.RawYearData{
	{Year: 2019, DepositRate: 7.0, InflationRate: 15.8, GDPGrowth: 9.0, ExchangeRate: 29.0, ForeignReserves: 3.0, GovtDebtGDP: 57.2},
	{Year: 2020, DepositRate: 7.0, InflationRate: 20.4, GDPGrowth: 6.1, ExchangeRate: 31.5, ForeignReserves: 2.8, GovtDebtGDP: 58.1},
	{Year: 2021, DepositRate: 7.0, InflationRate: 26.7, GDPGrowth: 6.3, ExchangeRate: 35.5, ForeignReserves: 2.5, GovtDebtGDP: 56.5},
	{Year: 2022, DepositRate: 7.0, InflationRate: 33.6, GDPGrowth: 5.4, ExchangeRate: 50.0, ForeignReserves: 2.1, GovtDebtGDP: 55.8},
	{Year: 2023, DepositRate: 7.0, InflationRate: 28.2, GDPGrowth: 6.0, ExchangeRate: 55.0, ForeignReserves: 2.3, GovtDebtGDP: 54.3},
	{Year: 2024, DepositRate: 7.0, InflationRate: 25.0, GDPGrowth: 6.5, ExchangeRate: 58.0, ForeignReserves: 2.6, GovtDebtGDP: 53.0},
}

func (db *DB) GetEconomicIndicators() (domain.APIResponse, error) {
	var response domain.APIResponse
	response.Metadata = metadata

	for _, yearData := range rawData {
		realInterestRate := yearData.DepositRate - yearData.InflationRate
		derivedData := domain.DerivedYearData{
			Year:             yearData.Year,
			RealInterestRate: round(realInterestRate, 1),
			GDPGrowth:        yearData.GDPGrowth,
			ExchangeRate:     yearData.ExchangeRate,
			ForeignReserves:  yearData.ForeignReserves,
			GovtDebtGDP:      yearData.GovtDebtGDP,
		}
		response.Data = append(response.Data, derivedData)
	}

	response.Source = "National Bank of Ethiopia"
	return response, nil
}

func round(val float64, precision int) float64 {
	p := 1.0
	for i := 0; i < precision; i++ {
		p *= 10
	}
	return float64(int(val*p+0.5)) / p
}
