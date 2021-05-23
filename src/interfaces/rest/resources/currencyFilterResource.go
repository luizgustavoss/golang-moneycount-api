package resources

import "moneycount-api/src/domain/model"

type CurrencyFilterResource struct {
	CurrencyCode string          `json:"currency_code"`
	Values       map[string]bool `json:"values"`
}

// NewCurrencyFilterResource creates a new CurrencyFilterResource based on CurrencyFilter model
func NewCurrencyFilterResource(currencyFilter model.CurrencyFilter) CurrencyFilterResource {

	var resource CurrencyFilterResource

	resource.CurrencyCode = currencyFilter.CurrencyCode
	resource.Values = currencyFilter.Values

	return resource
}

// NewCurrencyFilterFromCurrencyFilterResource creates a new CurrencyFilter model based on CurrencyFilterResource
func NewCurrencyFilterFromCurrencyFilterResource(resource CurrencyFilterResource) model.CurrencyFilter {

	var currencyFilter model.CurrencyFilter

	currencyFilter.CurrencyCode = resource.CurrencyCode
	currencyFilter.Values = resource.Values

	return currencyFilter
}