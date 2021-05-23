package routes

import (
	"moneycount-api/src/interfaces/rest/controllers"
	"net/http"
)

var currencyFiltersRoutes = []Route{
	{
		URI:                    "/currency-filters",
		Method:                 http.MethodGet,
		Function:               controllers.GetCurrencyFilter,
		RequiresAuthentication: false,
	},
}
