package router

import (
	"github.com/gorilla/mux"
	"moneycount-api/src/interfaces/rest/router/routes"
)

// Generates a router with configured routes
func Generate() *mux.Router {
	router := mux.NewRouter()
	return routes.ConfigureRoutes(router)
}