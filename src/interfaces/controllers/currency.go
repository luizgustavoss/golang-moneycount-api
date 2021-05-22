package controllers

import (
	"github.com/gorilla/mux"
	"moneycount-api/src/application/services"
	"moneycount-api/src/responses"
	"net/http"
)

// Get all available currencies
func GetAllCurrencies(w http.ResponseWriter, r *http.Request){

	currencies, error := services.ListCurrencies()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusOK, currencies)
}

// Get a specific currency
func GetCurrencyByCode(w http.ResponseWriter, r *http.Request){

	pathParameters := mux.Vars(r)
	code := pathParameters["code"]

	currency, error := services.GetCurrencyByCode(code)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusOK, currency)

}