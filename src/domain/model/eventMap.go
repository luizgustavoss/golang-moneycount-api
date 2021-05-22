package model

type EventMap struct {
	CurrencyCode string          `json:"currency_code"`
	Code         string          `json:"code"`
	Description  string          `json:"description"`
	CurrencyMap  CurrencyMap     `json:"currency_map"`
	Entries      []EventMapEntry `json:"entries"`
}
