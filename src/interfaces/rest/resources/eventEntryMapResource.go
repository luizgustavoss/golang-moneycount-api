package resources

import "moneycount-api/src/domain/model"

type EventEntryMapResource struct {
	CurrencyCode string              `json:"currency_code"`
	Code         string              `json:"code"`
	Description  string              `json:"description"`
	Value        float64             `json:"value"`
	CurrencyMap  CurrencyMapResource `json:"currency_map"`
}

// NewEventEntryMapResource creates a new EventEntryMapResource based on EventEntryMap model
func NewEventEntryMapResource(eventEntryMap model.EventEntryMap) EventEntryMapResource {

	var resource EventEntryMapResource

	resource.CurrencyCode = eventEntryMap.CurrencyCode
	resource.Code = eventEntryMap.Code
	resource.Description = eventEntryMap.Code
	resource.Value = eventEntryMap.Value
	resource.CurrencyMap = NewCurrencyMapResource(eventEntryMap.CurrencyMap)

	return resource
}