package resources

import "moneycount-api/src/domain/model"

type EventResource struct {
	Code        string                 `json:"code"`
	Description string                 `json:"description"`
	Entries     []EventEntryResource   `json:"entries"`
}

// NewEventFromEventResource creates a new Event model based on EventResource
func NewEventFromEventResource(resource EventResource) model.Event {

	var event model.Event

	event.Code = resource.Code
	event.Description = resource.Description
	event.Entries = make([]model.EventEntry, len(resource.Entries))

	for index, entry := range resource.Entries {
		event.Entries[index] = NewEventEntryFromEventEntryResource(entry)
	}

	return event
}