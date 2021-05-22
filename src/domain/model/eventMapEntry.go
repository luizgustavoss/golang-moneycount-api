package model

type EventMapEntry struct {
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Value       float64     `json:"value"`
	CurrencyMap CurrencyMap `json:"currency_map"`
}