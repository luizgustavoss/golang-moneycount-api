package routes

import (
	"moneycount-api/src/interfaces/rest/controllers"
	"net/http"
)

var currencyRoutes = []Route{
	{
		URI:                    "/v1/currencies",
		Method:                 http.MethodGet,
		Function:               controllers.GetAllCurrencies,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/v1/currencies/{code}",
		Method:                 http.MethodGet,
		Function:               controllers.GetCurrencyByCode,
		RequiresAuthentication: false,
	},
}
