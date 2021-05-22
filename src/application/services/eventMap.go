package services

import (
	"moneycount-api/src/application/commands"
	"moneycount-api/src/domain/model"
)

// CreateEventMap creates an EventMap based on provided Event info
func CreateEventMap(command commands.EventCommand) (model.EventMap, error){

	var eventMap model.EventMap
	var currency, error = GetCurrencyByCode(command.Filter.CurrencyCode)
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
