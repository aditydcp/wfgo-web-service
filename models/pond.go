package models

type Pond struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	FarmId string `json:"farm_id"`
}
