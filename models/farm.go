package models

type Farm struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	PondIds  []string  `json:"ponds"`
}