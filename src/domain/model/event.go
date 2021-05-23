package model

type Event struct {
	Code        string
	Description string
	Entries     []EventEntry
}
