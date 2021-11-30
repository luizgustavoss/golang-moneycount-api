package resources

import "moneycount-api/src/domain/model"

type EventMapEntryResource struct {
	Code        string              `json:"code"`
	Description string              `json:"description"`
	Value       float64             `json:"value"`
	CurrencyMap CurrencyMapResource `json:"currency_map"`
}

// NewEventMapEntryResource creates a new EventMapEntryResource based on EventMapEntry model
func NewEventMapEntryResource(eventMapEntry model.EventMapEntry) EventMapEntryResource {

	var resource EventMapEntryResource

	resource.Code = eventMapEntry.Code
	resource.Description = eventMapEntry.Description
	resource.Value = eventMapEntry.Value
	resource.CurrencyMap = NewCurrencyMapResource(eventMapEntry.CurrencyMap)

	return resource
}
