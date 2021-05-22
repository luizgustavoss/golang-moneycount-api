package commands

import "moneycount-api/src/domain/model"

type EventCommand struct {
	Event  model.Event          `json:"event"`
	Filter model.CurrencyFilter `json:"filter"`
}
