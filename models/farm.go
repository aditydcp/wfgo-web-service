package models

type Farm struct {
	ID      uint     `json:"id" gorm:"primary_key"`
	Name    string   `json:"name"`
	PondIds []string `json:"pond_ids"`
}
