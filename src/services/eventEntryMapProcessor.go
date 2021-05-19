package services

import (
	"moneycount-api/src/model"
	"moneycount-api/src/persistence"
)



// CreateEventEntryMap creates an EventEntryMap based on provided EventEntry info
func CreateEventEntryMap(command model.EventEntryCommand) (model.EventEntryMap, error){

	var eventEntryMap model.EventEntryMap
	var currency, error = persistence.NewCurrencyRepository().GetCurrencyByCode(command.Filter.CurrencyCode)
	if error != nil {
		return eventEntryMap, error
	}

	eventEntryMap.CurrencyCode = currency.Code
	eventEntryMap.Code = command.EventEntry.Code
	eventEntryMap.Description = command.EventEntry.Description
	eventEntryMap.Value = command.EventEntry.Value
	eventEntryMap.CurrencyMap = model.NewCurrencyMap()
	eventEntryMap.CurrencyMap.Build(command.EventEntry.Value, currency, command.Filter)

	return eventEntryMap, nil
}


// CreateEventMap creates an EventMap based on provided Event info
func CreateEventMap(command model.EventCommand) (model.EventMap, error){

	var eventMap model.EventMap
	var currency, error = persistence.NewCurrencyRepository().GetCurrencyByCode(command.Filter.CurrencyCode)
	if error != nil {
		return eventMap, error
	}

	eventMap.CurrencyCode = currency.Code
	eventMap.Code = command.Event.Code
	eventMap.Description = command.Event.Description
	eventMap.CurrencyMap = model.NewCurrencyMap()
	eventMap.Entries = make([]model.EventMapEntry, len(command.Event.Entries))

	for i, e := range command.Event.Entries {

		var eventMapEntry model.EventMapEntry
		eventMapEntry.Code = e.Code
		eventMapEntry.Description = e.Description
		eventMapEntry.Value = e.Value
		eventMapEntry.CurrencyMap = model.NewCurrencyMap()
		eventMapEntry.CurrencyMap.Build(e.Value, currency, command.Filter)

		eventMap.CurrencyMap.Add(eventMapEntry.CurrencyMap)
		eventMap.Entries[i] = eventMapEntry
	}

	return eventMap, nil
}