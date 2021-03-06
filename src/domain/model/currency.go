package model

import "errors"

type Currency struct {
	Code   string
	Name   string
	Symbol string
	Values []float64
}

// Validate check if currency has all required attributes
func (currency *Currency) Validate() error {

	if currency.Code == "" {
		return errors.New("invalid currency code")
	}

	if currency.Name == "" {
		return errors.New("invalid currency name")
	}

	if currency.Symbol == "" {
		return errors.New("invalid currency symbol")
	}

	if len(currency.Values) <= 0 {
		return errors.New("invalid currency values")
	}

	return nil
}
