package routes

import (
	"log"
	"net/http"
	"task/controllers/handlers"
	"task/controllers/repository"
	"task/db"
	"task/usecases"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Router struct {
	route *mux.Router
}

func NewRouter(r *mux.Router) *Router {
	return &Router{route: r}
}

func (r *Router) RegisterRoute() {
	// Stock endpoint
	db := db.NewRepository()
	stockRepository := repository.NewStockRepository(db)
	stockUsecase := usecases.NewStockUseCase(stockRepository)
	stockHandler := handlers.NewStockHandler(stockUsecase)
	stockRoute := r.route.PathPrefix("/api/v1").Subrouter()
	stockRoute.Handle("/stocks", http.HandlerFunc(stockHandler.GetAllStocks)).Methods("GET")
	// Bond endpoint
	bondRepository := repository.NewBondRepository(db)
	bondUsecase := usecases.NewBondUseCase(bondRepository)
	bondHandler := handlers.NewBondHandler(bondUsecase)
	bondRoute := r.route.PathPrefix("/api/v1").Subrouter()
	bondRoute.Handle("/bonds/search", http.HandlerFunc(bondHandler.GetBonds)).Methods("GET")

	// CMSP endpoint
	cmspRepository := repository.NewCMSPRepository(db)
	cmspUsecase := usecases.NewCMSPUseCase(cmspRepository)
	cmspHandler := handlers.NewCMSPHandler(*cmspUsecase)
	cmspRoute := r.route.PathPrefix("/api/v1").Subrouter()
	cmspRoute.Handle("/cmsp", http.HandlerFunc(cmspHandler.GetAllCMSPs)).Methods("GET")
	cmspRoute.Handle("/cmsp/{id:[0-9]+}", http.HandlerFunc(cmspHandler.GetCMSPByID)).Methods("GET")
}

func (r *Router) Run(addr string) error {
	// CORS configuration to allow all origins
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                       // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		AllowedHeaders: []string{"Content-Type", "Authorization"},           // Allowed headers
	})

	// Wrap the mux router with CORS middleware
	handler := corsHandler.Handler(r.route)

	// Run the server with CORS enabled
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, handler)
}
