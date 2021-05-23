package resources

type EventCommandResource struct {
	Event  EventResource          `json:"event"`
	Filter CurrencyFilterResource `json:"filter"`
}
