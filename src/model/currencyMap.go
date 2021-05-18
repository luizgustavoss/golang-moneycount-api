package model

type CurrencyMap struct {
	Map            map[string]int32 `json:"map"`
	RemainingValue float64          `json:"remaining_value"`
	TotalValue     float64          `json:"total_value"`
}
