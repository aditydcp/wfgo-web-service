package models

import "github.com/jinzhu/gorm"

type Farm struct {
	gorm.Model
	Id      string
	Name    string
	PondIds []string
}
