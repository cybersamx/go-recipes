package models

type State struct {
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
	Population   int    `json:"population"`
}
