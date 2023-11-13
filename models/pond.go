package models

import "github.com/jinzhu/gorm"

type Pond struct {
	gorm.Model
	Id     string
	Name   string
	FarmId string
}
