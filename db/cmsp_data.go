package db

import "task/domain"

var simulateCMSP = []domain.CMSP{
	{
		ID:           1,
		Name:         "CBE Capital S.C.",
		Type:         "Investment Bank",
		LicensedDate: "2025-03-21",
		Description:  "A newly licensed investment bank majority‑owned by Commercial Bank of Ethiopia, offering underwriting, capital raising, M&A, and advisory services.",
		Plan:         "Expand into portfolio management and Islamic finance.",
		Source:       "https://capitalethiopia.com/2025/03/24/ecma-licenses-five-new-capital-market-service-providers/",
	},
	{
		ID:           2,
		Name:         "Wegagen Capital Investment Bank S.C.",
		Type:         "Investment Bank",
		LicensedDate: "2025-03-21",
		Description:  "First private-sector investment bank under Wegagen Bank, providing capital raising, advisory and underwriting services.",
		Plan:         "Strengthen market outreach through parent bank network.",
		Source:       "https://www.capitalmarketethiopia.com/category/exchange/licensed-advisors/",
	},
	{
		ID:           3,
		Name:         "Ethio‑Fidelity Securities S.C.",
		Type:         "Securities Dealer",
		LicensedDate: "2025-03-21",
		Description:  "Ethiopia’s first licensed securities dealer to facilitate trading and improve market liquidity.",
		Plan:         "Launch electronic trading platform and broker training programs.",
		Source:       "https://www.stockmarket.et/ethiopian-capital-market-authority-licenses-five-new-service-providers/",
	},
	{
		ID:           4,
		Name:         "HST Investment Advisory Services PLC",
		Type:         "Securities Investment Advisor",
		LicensedDate: "2025-03-21",
		Description:  "Licensed investment adviser offering corporate and institutional investor guidance.",
		Plan:         "Offer ESG-compliant advisory and SME-focused financial planning.",
		Source:       "https://www.aksion.et/ethiopian-capital-market-ethio-fidelity-securities-new-licenses/",
	},
	{
		ID:           5,
		Name:         "Equation Securities Investment Advisor PLC",
		Type:         "Securities Investment Advisor",
		LicensedDate: "2025-03-21",
		Description:  "Investment advisory firm licensed by ECMA for individual and institutional clients.",
		Plan:         "Develop digital advisory platform and investor education programs.",
		Source:       "https://capitalethiopia.com/2025/03/24/ecma-licenses-five-new-capital-market-service-providers/",
	},
	{
		ID:           6,
		Name:         "Zemen Capital",
		Type:         "Investment Bank",
		LicensedDate: "2025-04-10",
		Description:  "A division of Zemen Bank focusing on corporate finance, investment research, and capital raising.",
		Plan:         "Build a bond issuance advisory service for large-scale infrastructure projects.",
		Source:       "https://www.zemenbank.com",
	},
	{
		ID:           7,
		Name:         "PrimePoint Capital",
		Type:         "Securities Investment Advisor",
		LicensedDate: "2025-05-02",
		Description:  "Privately held firm providing personal and institutional investment advice in emerging markets.",
		Plan:         "Launch mobile-based retail investor advisory tools.",
		Source:       "https://www.capitalmarketethiopia.com/primepoint-capital/",
	},
	{
		ID:           8,
		Name:         "Blue Nile Securities",
		Type:         "Securities Dealer",
		LicensedDate: "2025-05-10",
		Description:  "A licensed securities dealer committed to transparent execution and market access.",
		Plan:         "Open branch offices in five regional cities for retail investor access.",
		Source:       "https://www.bluenilesecurities.et/about",
	},
	{
		ID:           9,
		Name:         "Hibret Capital Partners",
		Type:         "Investment Bank",
		LicensedDate: "2025-06-01",
		Description:  "An investment bank focused on SMEs and early-stage startups in Ethiopia’s growing tech ecosystem.",
		Plan:         "Partner with fintech incubators to develop funding pipelines.",
		Source:       "https://www.hibretbank.com",
	},
	{
		ID:           10,
		Name:         "Lemat Investment Advisory",
		Type:         "Securities Investment Advisor",
		LicensedDate: "2025-06-15",
		Description:  "Advisory firm supporting cooperative unions, pensions, and rural development programs.",
		Plan:         "Launch investor literacy programs in Amharic and Afaan Oromo.",
		Source:       "Lemat internal press release, June 2025",
	},
}

func (db *DB) GetAllCMSPs() []domain.CMSP {
	return simulateCMSP
}
