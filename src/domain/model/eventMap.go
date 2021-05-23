package model

type EventMap struct {
	CurrencyCode string
	Code         string
	Description  string
	CurrencyMap  CurrencyMap
	Entries      []EventMapEntry
}
