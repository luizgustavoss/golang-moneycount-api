package routes

import (
	"moneycount-api/src/controllers"
	"net/http"
)

var currencyRoutes = []Route{
	{
		URI: "/currencies",
		Method: http.MethodGet,
		Function: controllers.GetAllCurrencies,
		RequiresAuthentication: false,
	},
	{
		URI: "/currencies/{code}",
		Method: http.MethodGet,
		Function: controllers.GetCurrencyByCode,
		RequiresAuthentication: false,
	},
}
