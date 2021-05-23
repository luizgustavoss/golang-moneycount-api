package services

import (
	"moneycount-api/src/domain/model"
)

// CreateEventMap creates an EventMap based on provided Event info
func CreateEventMap(event model.Event, currencyFilter model.CurrencyFilter) (model.EventMap, error) {

	var eventMap model.EventMap
	var currency, error = GetCurrencyByCode(currencyFilter.CurrencyCode)
	if error != nil {
		return eventMap, error
	}

	eventMap.CurrencyCode = currency.Code
	eventMap.Code = event.Code
	eventMap.Description = event.Description
	eventMap.CurrencyMap = model.NewCurrencyMap()
	eventMap.Entries = make([]model.EventMapEntry, len(event.Entries))

	for i, e := range event.Entries {

		var eventMapEntry model.EventMapEntry
		eventMapEntry.Code = e.Code
		eventMapEntry.Description = e.Description
		eventMapEntry.Value = e.Value
		eventMapEntry.CurrencyMap = model.NewCurrencyMap()
		eventMapEntry.CurrencyMap.Build(e.Value, currency, currencyFilter)

		eventMap.CurrencyMap.Add(eventMapEntry.CurrencyMap)
		eventMap.Entries[i] = eventMapEntry
	}

	return eventMap, nil
}
