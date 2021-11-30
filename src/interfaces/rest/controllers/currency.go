package controllers

import (
	"github.com/gorilla/mux"
	"moneycount-api/src/application/services"
	"moneycount-api/src/interfaces/rest/resources"
	"moneycount-api/src/interfaces/rest/responses"
	"net/http"
)

// GetAllCurrencies get all available CurrencyResource
func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {

	currencies, error := services.ListCurrencies()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusOK, resources.NewCurrencyResourceList(currencies))
}

// GetCurrencyByCode get a specific CurrencyResource
func GetCurrencyByCode(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	code := pathParameters["code"]

	currency, error := services.GetCurrencyByCode(code)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusOK, resources.NewCurrencyResource(currency))

}
