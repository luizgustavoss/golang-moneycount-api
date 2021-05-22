package controllers

import (
	"encoding/json"
	"io/ioutil"
	"moneycount-api/src/application/services"
	"moneycount-api/src/domain/model"
	"moneycount-api/src/responses"
	"net/http"
)

// CreateEventMap creates a currency map for the event based on a filter
func CreateEventMap(w http.ResponseWriter, r *http.Request){

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var command model.EventCommand
	if error = json.Unmarshal(requestBody, &command); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	eventMap, error := services.CreateEventMap(command)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	responses.JsonResponse(w, http.StatusCreated, eventMap)
}