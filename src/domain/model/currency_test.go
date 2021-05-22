package model

import "testing"

func TestCurrency(t *testing.T) {

	t.Run("should identify invalid code", func(t *testing.T){
		currency := Currency{Name: "Real", Symbol: "R$", Values: []float64{0.05, 0.10}}
		expected := currency.Validate()

		if expected.Error() != "invalid currency code" {
			t.Errorf("Expected: 'invalid currency code' | Returned: %s ", expected.Error())
		}
	})

	t.Run("should identify invalid name", func(t *testing.T){
		currency := Currency{Code: "BRL", Symbol: "R$", Values: []float64{0.05, 0.10}}
		expected := currency.Validate()

		if expected.Error() != "invalid currency name" {
			t.Errorf("Expected: 'invalid currency name' | Returned: %s ", expected.Error())
		}
	})

	t.Run("should identify invalid symbol", func(t *testing.T){
		currency := Currency{Code: "BRL", Name: "Real", Values: []float64{0.05, 0.10}}
		expected := currency.Validate()

		if expected.Error() != "invalid currency symbol" {
			t.Errorf("Expected: 'invalid currency symbol' | Returned: %s ", expected.Error())
		}
	})

	t.Run("should identify invalid values", func(t *testing.T){
		currency := Currency{Code: "BRL", Name: "Real", Symbol: "R$"}
		expected := currency.Validate()

		if expected.Error() != "invalid currency values" {
			t.Errorf("Expected: 'invalid currency values' | Returned: %s ", expected.Error())
		}
	})

}
