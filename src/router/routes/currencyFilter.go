package routes

import (
	"moneycount-api/src/controllers"
	"net/http"
)

var currencyFilterRoutes = []Route{
	{
		URI: "/currency-filters",
		Method: http.MethodGet,
		Function: controllers.GetCurrencyFilter,
		RequiresAuthentication: false,
	},
}
