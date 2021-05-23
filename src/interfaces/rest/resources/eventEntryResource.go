package resources

import "moneycount-api/src/domain/model"

type EventEntryResource struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

// NewEventEntryFromEventEntryResource creates a new EventEntry model based on EventEntryResource
func NewEventEntryFromEventEntryResource(resource EventEntryResource) model.EventEntry {

	var eventEntry model.EventEntry

	eventEntry.Code = resource.Code
	eventEntry.Description = resource.Description
	eventEntry.Value = resource.Value

	return eventEntry
}