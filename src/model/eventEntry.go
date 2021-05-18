package model

type EventEntry struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}
