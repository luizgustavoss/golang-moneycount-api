package resources

import "moneycount-api/src/domain/model"

type CurrencyMapResource struct {
	Map            map[string]int32 `json:"map"`
	RemainingValue float64          `json:"remaining_value"`
	TotalValue     float64          `json:"total_value"`
}

// NewCurrencyMapResource creates a new CurrencyMapResource based on CurrencyMap model
func NewCurrencyMapResource(currencyMap model.CurrencyMap) CurrencyMapResource {

	var resource CurrencyMapResource

	resource.Map = currencyMap.Map
	resource.RemainingValue = currencyMap.RemainingValue
	resource.TotalValue = currencyMap.TotalValue

	return resource
}
