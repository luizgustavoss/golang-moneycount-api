package model

import (
	"fmt"
)

type CurrencyFilter struct {
	Values map[string]bool `json:"values"`
}

func NewCurrencyFilter(currency Currency) CurrencyFilter {

	var values = make(map[string]bool)

	for _, v := range currency.Values{
		values[fmt.Sprintf("%.2f", v)] = true
	}

	return CurrencyFilter{values}
}