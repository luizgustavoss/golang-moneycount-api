package services

import "io/ioutil"

func ReadCurrencyFileContent() string {
	fileData, _ := ioutil.ReadFile("resources/currencies.json")
	return string(fileData)
}
