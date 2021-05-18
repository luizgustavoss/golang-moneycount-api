package model

type EventCommand struct {
	CurrencyCode string         `json:"currency_code"`
	Event        Event          `json:"event"`
	Filter       CurrencyFilter `json:"filter"`
}
