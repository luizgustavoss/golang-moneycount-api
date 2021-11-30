package controllers

import (
	"encoding/json"
	"io/ioutil"
	"moneycount-api/src/application/services"
	"moneycount-api/src/interfaces/rest/resources"
	"moneycount-api/src/interfaces/rest/responses"
	"net/http"
)

// CreateEventMap creates a currency map for the event based on a filter
func CreateEventMap(w http.ResponseWriter, r *http.Request) {

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var resource resources.EventCommandResource
	if error = json.Unmarshal(requestBody, &resource); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	event := resources.NewEventFromEventResource(resource.Event)
	currencyFilter := resources.NewCurrencyFilterFromCurrencyFilterResource(resource.Filter)

	eventMap, error := services.CreateEventMap(event, currencyFilter)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	responses.JsonResponse(w, http.StatusCreated, resources.NewEventMapResource(eventMap))
}
