package resources

import "moneycount-api/src/domain/model"

type CurrencyMapResource struct {
	Currencies     []CurrencyMapItemResource `json:"currencies"`
	RemainingValue float64                   `json:"remaining_value"`
	TotalValue     float64                   `json:"total_value"`
}

type CurrencyMapItemResource struct {
	CurrencyValue string `json:"currency_value"`
	Count         int32  `json:"count"`
}

// NewCurrencyMapResource creates a new CurrencyMapResource based on CurrencyMap model
func NewCurrencyMapResource(currencyMap model.CurrencyMap) CurrencyMapResource {

	var resource CurrencyMapResource

	resource.RemainingValue = currencyMap.RemainingValue
	resource.TotalValue = currencyMap.TotalValue

	for key, value := range currencyMap.Map {
		var item CurrencyMapItemResource
		item.CurrencyValue = key
		item.Count = value
		resource.Currencies = append(resource.Currencies, item)
	}

	return resource
}
