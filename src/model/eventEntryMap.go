package model

type EventEntryMap struct {
	CurrencyCode string      `json:"currency_code"`
	Code         string      `json:"code"`
	Description  string      `json:"description"`
	Value        float64     `json:"value"`
	CurrencyMap  CurrencyMap `json:"currency_map"`
}
