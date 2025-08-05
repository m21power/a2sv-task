# Capital Market Data API

This project is a Go-based RESTful API for retrieving and filtering data about stocks, bonds, and Capital Market Service Providers (CMSPs) in Ethiopia. It is designed for educational and demonstration purposes, simulating real-world financial data and endpoints.

## Features

- **Stocks**: Retrieve all stocks, filter by gainers/losers.
- **Bonds**: Search bonds by coupon rate and maturity date.
- **CMSPs**: List all CMSPs, filter by type or license date, and get details by ID.
- **CORS enabled**: Ready for frontend integration.

## Project Structure

```
.
├── main.go                # Entry point, starts the HTTP server
├── routes/                # API route registration and server setup
├── controllers/
│   └── handlers/          # HTTP handlers for stocks, bonds, CMSPs
│   └── repository/        # Data access layer (repositories)
├── db/                    # Simulated data sources
├── domain/                # Core business models and interfaces
├── usecases/              # Business logic for each entity
├── types/                 # Shared type definitions
├── utils/                 # Utility functions (error handling, JSON, etc.)
```

## API Endpoints

- **GET `/api/v1/stocks`**  
  List all stocks.  
  Query param: `type=gainer|loser` (optional)

- **GET `/api/v1/bonds/search`**  
  Search bonds.  
  Query params: `minCoupon`, `maxCoupon`, `maturityBefore`, `maturityAfter` (all optional, dates in `YYYY-MM-DD`)

- **GET `/api/v1/cmsp`**  
  List all CMSPs.  
  Query params: `type`, `licensedBefore`, `licensedAfter` (dates in `YYYY-MM-DD`)

- **GET `/api/v1/cmsp/{id}`**  
  Get CMSP by ID.

## How to Run

1. **Install Go** (if not already):  
   https://golang.org/doc/install

2. **Clone the repository** and navigate to the project directory.

3. **Run the server:**
   ```bash
   go run main.go
   ```
   The API will be available at `http://localhost:8080`.

## Example Requests

- List all stocks:  
  `GET http://localhost:8080/api/v1/stocks`

- List bond search:  
  `GET http://localhost:8080/api/v1/bonds/search?minCoupon=4.5&maturityBefore=2028-01-01`

- List all CMSPs:  
  `GET http://localhost:8080/api/v1/cmsp?type=Investment%20Bank`

## Notes

- All data is simulated and for demonstration only.
- CORS is enabled for all origins.
