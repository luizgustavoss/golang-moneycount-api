package model

type EventEntryCommand struct {
	EventEntry   EventEntry     `json:"event_entry"`
	Filter       CurrencyFilter `json:"filter"`
}