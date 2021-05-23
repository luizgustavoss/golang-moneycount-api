package resources

import "moneycount-api/src/domain/model"

type CurrencyResource struct {
	Code   string    `json:"code"`
	Name   string    `json:"name"`
	Symbol string    `json:"symbol"`
}

// NewCurrencyResource creates a new CurrencyResource based on Currency model
func NewCurrencyResource(currency model.Currency) CurrencyResource {

	var resource CurrencyResource

	resource.Code = currency.Code
	resource.Symbol = currency.Symbol
	resource.Name = currency.Name

	return resource
}

// NewCurrencyResourceList creates a new CurrencyResource list based on Currency model list
func NewCurrencyResourceList(currencies []model.Currency) []CurrencyResource {

	resourceList := make([]CurrencyResource, len(currencies))

	for index, currency := range currencies {
		var resource CurrencyResource

		resource.Code = currency.Code
		resource.Symbol = currency.Symbol
		resource.Name = currency.Name

		resourceList[index] = resource
	}
	return resourceList
}