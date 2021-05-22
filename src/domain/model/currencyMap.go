package model

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type CurrencyMap struct {
	Map            map[string]int32 `json:"map"`
	RemainingValue float64          `json:"remaining_value"`
	TotalValue     float64          `json:"total_value"`
}

func (c *CurrencyMap) Build(value float64, currency Currency, filter CurrencyFilter) error {

	if currency.Code != filter.CurrencyCode{
		return errors.New("invalid currency for filter")
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(currency.Values)))

	c.TotalValue = value
	c.RemainingValue = value

	for _, v := range currency.Values {
		key := fmt.Sprintf("%.2f", v)

		c.Map[key] = 0

		if filter.Values[key] {
			count := c.RemainingValue / v
			c.Map[key] = int32(math.Trunc(count))
			c.RemainingValue = math.Mod(c.RemainingValue, v)
		}
	}

	if r, error := strconv.ParseFloat(fmt.Sprintf("%.2f", c.RemainingValue), 64); error == nil {
		c.RemainingValue = r
	} else{
		return error
	}

	return nil
}

func (c *CurrencyMap) Add(currencyMap CurrencyMap) {
	for k, v := range currencyMap.Map {
		c.Map[k] += v
	}

	c.RemainingValue += currencyMap.RemainingValue
	c.TotalValue += currencyMap.TotalValue
}


func NewCurrencyMap() CurrencyMap {
	return CurrencyMap{make(map[string]int32), 0, 0}
}