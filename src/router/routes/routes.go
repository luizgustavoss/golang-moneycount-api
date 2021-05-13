package routes

import (
	"github.com/gorilla/mux"
	"moneycount-api/src/middlewares"
	"net/http"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {

	routes := currencyRoutes
	routes = append(routes, currencyFilterRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, middlewares.LogRequest(route.Function)).Methods(route.Method)
	}
	return r
}

