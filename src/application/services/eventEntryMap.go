package services

import (
	"moneycount-api/src/application/commands"
	"moneycount-api/src/domain/model"
)

// CreateEventEntryMap creates an EventEntryMap based on provided EventEntry info
func CreateEventEntryMap(command commands.EventEntryCommand) (model.EventEntryMap, error){

	var eventEntryMap model.EventEntryMap
	var currency, error = GetCurrencyByCode(command.Filter.CurrencyCode)
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

