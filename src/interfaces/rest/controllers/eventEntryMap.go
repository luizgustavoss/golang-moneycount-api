package controllers

import (
	"encoding/json"
	"io/ioutil"
	"moneycount-api/src/application/services"
	"moneycount-api/src/interfaces/rest/resources"
	"moneycount-api/src/interfaces/rest/responses"
	"net/http"
)

// CreateEventEntryMap creates a currency map for the event entry based on a filter
func CreateEventEntryMap(w http.ResponseWriter, r *http.Request) {

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var resource resources.EventEntryCommandResource
	if error = json.Unmarshal(requestBody, &resource); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	eventEntry := resources.NewEventEntryFromEventEntryResource(resource.EventEntry)
	currencyFilter := resources.NewCurrencyFilterFromCurrencyFilterResource(resource.Filter)

	eventEntryMap, error := services.CreateEventEntryMap(eventEntry, currencyFilter)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	responses.JsonResponse(w, http.StatusCreated, resources.NewEventEntryMapResource(eventEntryMap))
}
