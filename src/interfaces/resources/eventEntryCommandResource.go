package resources

type EventEntryCommandResource struct {
	EventEntry EventEntryResource     `json:"event_entry"`
	Filter     CurrencyFilterResource `json:"filter"`
}
