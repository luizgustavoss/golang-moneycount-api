package model

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type CurrencyMap struct {
	Map            map[string]int32
	RemainingValue float64
	TotalValue     float64
}

func (c *CurrencyMap) Build(value float64, currency Currency, filter CurrencyFilter) error {

	if currency.Code != filter.CurrencyCode {
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
			if r, error := round(math.Mod(c.RemainingValue, v)); error == nil {
				c.RemainingValue = r
			} else {
				return error
			}
		}
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

func round(value float64) (float64, error) {
	if value, error := strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64); error == nil {
		return value, nil
	} else {
		return value, error
	}
}
