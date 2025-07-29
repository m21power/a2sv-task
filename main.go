package main

import (
	"task/routes"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	r := routes.NewRouter(route)
	r.RegisterRoute()
	r.Run(":8080")
}
