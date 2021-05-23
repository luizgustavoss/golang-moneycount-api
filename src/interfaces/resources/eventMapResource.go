package resources

import "moneycount-api/src/domain/model"

type EventMapResource struct {
	CurrencyCode string                  `json:"currency_code"`
	Code         string                  `json:"code"`
	Description  string                  `json:"description"`
	CurrencyMap  CurrencyMapResource     `json:"currency_map"`
	Entries      []EventMapEntryResource `json:"entries"`
}

// NewEventMapResource creates a new EventMapResource based on EventMap model
func NewEventMapResource(eventMap model.EventMap) EventMapResource {

	var resource EventMapResource

	resource.CurrencyCode = eventMap.CurrencyCode
	resource.Code = eventMap.Code
	resource.Description = eventMap.Description
	resource.CurrencyMap = NewCurrencyMapResource(eventMap.CurrencyMap)
	resource.Entries = make([]EventMapEntryResource, len(eventMap.Entries))

	for index, eventMapEntry := range eventMap.Entries {
		resource.Entries[index] = NewEventMapEntryResource(eventMapEntry)
	}

	return resource
}