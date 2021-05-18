package model

type EventEntryCommand struct {
	CurrencyCode string         `json:"currency_code"`
	EventEntry   EventEntry     `json:"event_entry"`
	Filter       CurrencyFilter `json:"filter"`
}
