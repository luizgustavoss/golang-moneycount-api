package resources

import "moneycount-api/src/domain/model"

type CurrencyFilterResource struct {
	CurrencyCode string                       `json:"currency_code"`
	Values       []CurrencyFilterItemResource `json:"values"`
}

type CurrencyFilterItemResource struct {
	CurrencyValue string `json:"currency_value"`
	ShouldBeUsed  bool   `json:"should_use"`
}

// NewCurrencyFilterResource creates a new CurrencyFilterResource based on CurrencyFilter model
func NewCurrencyFilterResource(currencyFilter model.CurrencyFilter) CurrencyFilterResource {

	var resource CurrencyFilterResource

	resource.CurrencyCode = currencyFilter.CurrencyCode

	for key, value := range currencyFilter.Values {
		var item CurrencyFilterItemResource
		item.CurrencyValue = key
		item.ShouldBeUsed = value
		resource.Values = append(resource.Values, item)
	}

	return resource
}

// NewCurrencyFilterFromCurrencyFilterResource creates a new CurrencyFilter model based on CurrencyFilterResource
func NewCurrencyFilterFromCurrencyFilterResource(resource CurrencyFilterResource) model.CurrencyFilter {

	var currencyFilter model.CurrencyFilter

	currencyFilter.CurrencyCode = resource.CurrencyCode
	currencyFilter.Values = make(map[string]bool)

	for _, item := range resource.Values {
		currencyFilter.Values[item.CurrencyValue] = item.ShouldBeUsed
	}

	return currencyFilter
}
