package model

type EventEntryMap struct {
	CurrencyCode string
	Code         string
	Description  string
	Value        float64
	CurrencyMap  CurrencyMap
}
