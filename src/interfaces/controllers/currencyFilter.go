package controllers

import (
	"moneycount-api/src/application/services"
	"moneycount-api/src/interfaces/resources"
	"moneycount-api/src/interfaces/responses"
	"net/http"
	"strings"
)


// GetCurrencyFilter get a specific CurrencyResource
func GetCurrencyFilter(w http.ResponseWriter, r *http.Request){

	currencyCode := strings.ToUpper(r.URL.Query().Get("currency-code"))

	currency, error := services.GetCurrencyByCode(currencyCode)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	currencyFilter := services.NewCurrencyFilter(currency)

	responses.JsonResponse(w, http.StatusOK, resources.NewCurrencyFilterResource(currencyFilter))
}