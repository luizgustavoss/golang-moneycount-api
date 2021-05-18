package controllers

import (
	"encoding/json"
	"io/ioutil"
	"moneycount-api/src/model"
	"moneycount-api/src/responses"
	"moneycount-api/src/services"
	"net/http"
)

// CreateEventMap creates a currency map for the event based on a filter
func CreateEventMap(w http.ResponseWriter, r *http.Request){

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var event model.EventCommand
	if error = json.Unmarshal(requestBody, &event); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	eventMap, error := services.CreateEventMap(event)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	responses.JsonResponse(w, http.StatusCreated, eventMap)
}