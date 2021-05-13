package controllers

import (
	"moneycount-api/src/model"
	"moneycount-api/src/persistence"
	"moneycount-api/src/responses"
	"net/http"
	"strings"
)


// Get a specific currency
func GetCurrencyFilter(w http.ResponseWriter, r *http.Request){

	currencyCode := strings.ToUpper(r.URL.Query().Get("currency-code"))

	repository := persistence.NewCurrencyRepository()

	currency, error := repository.GetCurrencyByCode(currencyCode)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	currencyFilter := model.NewCurrencyFilter(currency)

	responses.JsonResponse(w, http.StatusOK, currencyFilter)

}