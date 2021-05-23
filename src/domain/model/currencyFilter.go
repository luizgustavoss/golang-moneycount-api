package model

type CurrencyFilter struct {
	CurrencyCode string
	Values       map[string]bool
}
