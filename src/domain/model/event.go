package model

type Event struct {
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Entries     []EventEntry `json:"entries"`
}
