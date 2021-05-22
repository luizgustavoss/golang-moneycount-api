package controllers

import (
	"moneycount-api/src/application/services"
	"moneycount-api/src/domain/model"
	"moneycount-api/src/responses"
	"net/http"
	"strings"
)


// Get a specific currency
func GetCurrencyFilter(w http.ResponseWriter, r *http.Request){

	currencyCode := strings.ToUpper(r.URL.Query().Get("currency-code"))

	currency, error := services.GetCurrencyByCode(currencyCode)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	currencyFilter := model.NewCurrencyFilter(currency)

	responses.JsonResponse(w, http.StatusOK, currencyFilter)
}