package services

import (
	"encoding/json"
	"moneycount-api/src/domain/model"
	"moneycount-api/src/infrastructure/services"
)

// ListCurrencies lists available currencies
func ListCurrencies() ([]model.Currency, error) {

	var currencies []model.Currency

	if error := json.Unmarshal([]byte(services.ReadCurrencyFileContent()), &currencies); error != nil {
		return nil, error
	}

	return currencies, nil
}

// GetCurrencyByCode get a currency by its code
func GetCurrencyByCode(code string) (model.Currency, error) {

	var currency model.Currency

	currencies, error := ListCurrencies()
	if error != nil{
		return currency, error
	}

	for _, c := range currencies {
		if c.Code == code{
			currency = c
			break
		}
	}
	return currency, nil
}

