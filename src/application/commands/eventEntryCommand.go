package commands

import "moneycount-api/src/domain/model"

type EventEntryCommand struct {
	EventEntry model.EventEntry     `json:"event_entry"`
	Filter     model.CurrencyFilter `json:"filter"`
}