package services

import (
	"fmt"
	"moneycount-api/src/domain/model"
)

func NewCurrencyFilter(currency model.Currency) model.CurrencyFilter {
	var values = make(map[string]bool)
	for _, v := range currency.Values {
		values[fmt.Sprintf("%.2f", v)] = true
	}
	return model.CurrencyFilter{CurrencyCode: currency.Code, Values: values}
}
