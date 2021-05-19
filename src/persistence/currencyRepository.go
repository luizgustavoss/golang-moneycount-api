package persistence

import (
	"encoding/json"
	"io/ioutil"
	"moneycount-api/src/model"
)

type CurrencyRepository struct {

}

// ListCurrencies lists available currencies
func (repository CurrencyRepository) ListCurrencies() ([]model.Currency, error) {

	var currencies []model.Currency

	if error := json.Unmarshal([]byte(readCurrencyFileContent()), &currencies); error != nil {
		return nil, error
	}

	return currencies, nil
}

// GetCurrencyByCode get a currency by its code
func (repository CurrencyRepository) GetCurrencyByCode(code string) (model.Currency, error) {

	var currency model.Currency

	currencies, error := repository.ListCurrencies()
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

// NewCurrencyRepository factory
func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}

func readCurrencyFileContent() string {
	fileData, _ := ioutil.ReadFile("resources/currencies.json")
	return string(fileData)
}