package model

type EventCommand struct {
	Event        Event          `json:"event"`
	Filter       CurrencyFilter `json:"filter"`
}
