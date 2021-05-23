package services

import (
	"moneycount-api/src/domain/model"
)

// CreateEventEntryMap creates an EventEntryMap based on provided EventEntry info
func CreateEventEntryMap(eventEntry model.EventEntry, currencyFilter model.CurrencyFilter) (model.EventEntryMap, error){

	var eventEntryMap model.EventEntryMap
	var currency, error = GetCurrencyByCode(currencyFilter.CurrencyCode)
	if error != nil {
		return eventEntryMap, error
	}

	eventEntryMap.CurrencyCode = currency.Code
	eventEntryMap.Code = eventEntry.Code
	eventEntryMap.Description = eventEntry.Description
	eventEntryMap.Value = eventEntry.Value
	eventEntryMap.CurrencyMap = model.NewCurrencyMap()
	eventEntryMap.CurrencyMap.Build(eventEntry.Value, currency, currencyFilter)

	return eventEntryMap, nil
}

